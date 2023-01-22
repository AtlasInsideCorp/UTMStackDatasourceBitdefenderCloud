package main

import (
	"os"
	"path"
)

func checkCerts() bool {
	//Check if exist key certificate
	if _, err := os.Stat(path.Join(getMyPath(), "certs", getenv("KEY"))); os.IsNotExist(err) {
		return false
	}
	//Check if exist certificate
	if _, err := os.Stat(path.Join(getMyPath(), "certs", getenv("CERT"))); os.IsNotExist(err) {
		return false
	}
	return true
}
