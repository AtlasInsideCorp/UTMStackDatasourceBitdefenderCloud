package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/gorilla/mux"
)

type JsonConf struct {
	Port       string `json:"port"`
	SyslogPort string `json:"syslog_port"`
	Protocol   string `json:"protocol"`
	Target     string `json:"target"`
	AuthHeader string `json:"authentication_string"`
}

type BodyEvents struct {
	CEF    string   `json:"cef"`
	Events []string `json:"events"`
}

var syslogHelper EpsSyslogHelper
var config JsonConf

func GetBDGZLogs(w http.ResponseWriter, r *http.Request) {
	var newBody BodyEvents

	//Check if the authorization exist
	if r.Header.Get("authorization") == "" {
		messag := "401 Missing Authorization Header"
		fmt.Println(messag)
		j, err := json.Marshal(messag)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(j)
	}
	//Check if the authorization is valid
	authorizationString := r.Header.Get("authorization")
	if config.AuthHeader != authorizationString {
		messag := "401 Invalid Authentication Credentials"
		fmt.Println(messag)
		j, err := json.Marshal(messag)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(j)
	}

	err := json.NewDecoder(r.Body).Decode(&newBody)
	if err != nil {
		panic(err)
	}

	events := newBody.Events
	syslogHelper.sentToSyslog(events)

	j, err := json.Marshal("HTTP 200 OK")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func serverUp(option int) {
	certPath := path.Join(getMyPath(), "certs")
	if option == 1 {
		certPath = "/usr/src/app/certs"
	}

	//Defining configuration
	config = JsonConf{
		Port:       ":" + getenv("CONNECTOR_PORT"),
		SyslogPort: getenv("SYSLOG_PORT"),
		Protocol:   getenv("SYSLOG_PROTOCOL"),
		Target:     getenv("SYSLOG_HOST"),
		AuthHeader: generateAuthCode(getenv("BDGZ_API_KEY")),
	}

	//Initialing connection with syslogServer
	syslogHelper.init(config)

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api", GetBDGZLogs).Methods("POST")

	server := &http.Server{
		Addr:           ":" + getenv("CONNECTOR_PORT"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Listening in port %s...\n", getenv("CONNECTOR_PORT"))
	err := server.ListenAndServeTLS(path.Join(certPath, getenv("CERT_NAME")), path.Join(certPath, getenv("KEY_NAME")))
	fmt.Println(err)

	//Close connection with syslogServer
	syslogHelper.clientSyslog.Close()
}
