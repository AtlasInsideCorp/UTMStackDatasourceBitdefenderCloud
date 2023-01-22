package main

import (
	"log"
)

func main() {
	//Loading environment variables
	loadEnv()

	//Check if certificates exist
	if !checkCerts() {
		log.Fatal("Error, certificates not found in path CURRENT_PATH/certs")
	}

	serverUp()

}
