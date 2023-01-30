# UTMStackDatasourceBitdefenderCloud
### Version 1.0.1

_UTMStackDatasourceBitdefenderCloud is a connector that receives logs from Bitdefender GravityZone Cloud and sends them to the UTMStack syslog_

_The connector uses the POST method to receive authenticated and protected messages from the GravityZone event push service. Parses the message and then forwards it to a UTMStack Syslog server_

_It has a set of tools that allow you to properly configure the Bitdefender GravityZone Cloud PUSH API_

## Getting started
_Before starting, you must generate an APIKEY in your Bitdefender control account. For it:_
* Login to ***https://cloud.gravityzone.bitdefender.com/***
* Click on your username in the top right corner of the console and choose MY ACCOUNT
* Select the APIs you want to use, in this case it will be just Event Push Service API
* Click on save, an ApiKey will be generated for the selected API

_This tool uses some configurations that you must define before starting, for this you must create an .env file in the root of this project, and define the following environment variables:_
* BDGZ_API_KEY= "APIKEY obtained in the previous process"
* PORT = "port that the connector will use to listen"
* SYSLOG_PORT = "port of the syslog server where the logs will be sent"
* PROTO="server protocol"
* TARGET = "IP address or domain of the syslog server where it will send the logs"
* KEY= "certificate key file name, in .key format"
* CERT= "security certificate name in .crt format"
* DEBUG= "defines if the server will be in debug mode. Possible values: true/false"
* ISCONTAINER= "defines if the connector will be raised in a docker container or not. Possible values: true/false"
* BDGZ_ACCESS_URL= "You can find this url in the My Account/Control Center API/Access URL tab of your account, and add /v1.0/jsonrpc/push to it, in most cases it is ***https://cloud.gravityzone.bitdefender .com/api/v1.0/jsonrpc/push***"
* BDGZ_URL= "url of the server or log collector that will receive the Bitdefender logs, example: https://example.com:8000"

_The security certificate and key must be located in the certs folder_

### Requirements

go 1.19
github.com/RackSec/srslog
github.com/gorilla/mux
github.com/joho/godotenv


## Deploy

_First you must build your executable, for this:_

```
go build
```

_Then run the program and the connector will be listening on the designated port_

```
.\gz_utmstackconn
```

_If you want to build this connector with Docker, just check the port in docker-compose.yml, by default it is set to port 8000 and define the environment variable in the .env ISCONTAINER= "true" . Then you can run:_

```
docker compose up
```

_Then, to send your configuration to the Bitdefender Push API, open another cmd console and run the program passing it the sentConfig parameter._

```
.\gz_utmstackconn sentConfig
```

_If you want to check your Bitdefender Push API configuration, you should run its executable passing it the getConfig parameter._

```
.\gz_utmstackconn getConfig
```

## Running tests

_Then, to send a test log to Bitdefender's Push API, you must run its executable passing it the logTest parameter._

```
.\gz_utmstackconn.exe logTest
```

_At this time, the test logs should be arriving at your server or log collector_

## License
_This project is under AGPL-3.0 license_
 