# UTMStackDatasourceBitdefenderCloud

_UTMStackDatasourceBitdefenderCloud is a connector that receives logs from Bitdefender GravityZone Cloud and sends them to the UTMStack syslog_

_The connector uses the POST method to receive authenticated and protected messages from the GravityZone event push service. Parses the message and then forwards it to a UTMStack Syslog server_

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

_Then it will only be to execute the program and that's it, the connector will be listening on the designated port_

## License
_This project is under AGPL-3.0 license_
 