package main

import (
	"os"
	"path"
)

func checkCerts() bool {
	var certPath string

	if getenv("ISCONTAINER") == "true" {
		certPath = "/usr/src/app/certs"
	} else {
		certPath = path.Join(getMyPath(), "certs")
	}
	//Check if exist key certificate
	if _, err := os.Stat(path.Join(certPath, getenv("KEY"))); os.IsNotExist(err) {
		return false
	}
	//Check if exist certificate
	if _, err := os.Stat(path.Join(certPath, getenv("CERT"))); os.IsNotExist(err) {
		return false
	}
	return true
}
