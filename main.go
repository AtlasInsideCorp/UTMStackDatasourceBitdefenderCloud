package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		arg := os.Args[1]
		switch arg {
		case "sentConfig":
			//Loading environment variables
			loadEnv()

			//Sending configuration
			fmt.Println("Sending configuration...")
			resp := setPushEventSettings()
			fmt.Println(resp.Status)
			myBody, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(myBody))
			defer resp.Body.Close()

		case "getConfig":
			//Loading environment variables
			loadEnv()

			fmt.Println("Getting configuration...")
			resp := getPushEventSettings()
			fmt.Println(resp.Status)
			myBody, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(myBody))
			defer resp.Body.Close()

		case "logTest":
			//Loading environment variables
			loadEnv()

			fmt.Println("Send Event Test...")
			resp := sendTestPushEvent()
			fmt.Println(resp.Status)
			myBody, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(myBody))
			defer resp.Body.Close()

		case "run":
			//Loading environment variables
			loadEnv()

			//Check if certificates exist
			if !checkCerts(2) {
				log.Fatal("Error, certificates not found in path CURRENT_PATH/certs")
			}

			serverUp(2)
		}
	} else {
		//Check if certificates exist
		if !checkCerts(1) {
			log.Fatal("Error, certificates not found in path CURRENT_PATH/certs")
		}

		serverUp(1)
	}

}
