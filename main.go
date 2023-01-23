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
		}
	} else {
		//Loading environment variables
		loadEnv()

		//Check if certificates exist
		if !checkCerts() {
			log.Fatal("Error, certificates not found in path CURRENT_PATH/certs")
		}

		serverUp()
	}

}
