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

func serverUp() {
	//Defining configuration
	config = JsonConf{
		Port:       getenv("PORT"),
		SyslogPort: getenv("SYSLOG_PORT"),
		Protocol:   getenv("PROTO"),
		Target:     getenv("TARGET"),
		AuthHeader: generateAuthCode(getenv("BDGZ_API_KEY")),
	}

	//Initialing connection with syslogServer
	syslogHelper.init(config)

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api", GetBDGZLogs).Methods("POST")

	server := &http.Server{
		Addr:           getenv("PORT"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Listening in port %s...\n", getenv("PORT"))
	err := server.ListenAndServeTLS(path.Join(getMyPath(), "certs", "server.crt"), path.Join(getMyPath(), "certs", "server.key"))
	fmt.Println(err)

	//Close connection with syslogServer
	syslogHelper.clientSyslog.Close()
}
