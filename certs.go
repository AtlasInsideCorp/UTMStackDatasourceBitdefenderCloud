package main

import (
	"os"
	"path"
)

func checkCerts(option int) bool {
	certPath := path.Join(getMyPath(), "certs")

	if option == 1 {
		certPath = "/usr/src/app/certs"
	}
	//Check if exist key certificate
	if _, err := os.Stat(path.Join(certPath, getenv("KEY_NAME"))); os.IsNotExist(err) {
		return false
	}

	//Check if exist certificate
	if _, err := os.Stat(path.Join(certPath, getenv("CERT_NAME"))); os.IsNotExist(err) {
		return false
	}

	return true
}
