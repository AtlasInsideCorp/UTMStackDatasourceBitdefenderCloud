# UTMStack Datasource for BitdefenderCloud
### Version 1.2.0

## Decription

UTMStack Datasource for BitdefenderCloud is a connector developed in golang that receives logs from Bitdefender GravityZone Cloud and sends them to syslog server

The connector uses the POST method to receive authenticated and protected messages from the GravityZone event push service. Parses the message and then forwards it to a UTMStack Syslog server

## Getting started

Before starting, you must generate an APIKEY in your Bitdefender control account. For it:
* Login to ***https://cloud.gravityzone.bitdefender.com/***
* Click on your username in the top right corner of the console and choose MY ACCOUNT
* Select the APIs you want to use, in this case it will be just Event Push Service API
* Click on save, an ApiKey will be generated for the selected API

## Usage

This tool uses some environment variables that you must define before starting:
* _BDGZ_API_KEY_: "APIKEY obtained in the previous process"
* _SYSLOG_PROTOCOL_: "protocol of the listening syslog server. Example: `tcp`"
* _SYSLOG_HOST_: "IP address or Host of the listening syslog server where it will send the logs."
* _SYSLOG_PORT_: "port of the listening syslog server where the logs will be sent. Example: `514`"
* _CONNECTOR_PORT_: "port that the connector will use to listen. Example: `8000`"
* _KEY_NAME_: "certificate key file name, in .key format. Example: `server.key`"
* _CERT_NAME_: "security certificate name in .crt format. Example: `server.crt`"
* _BDGZ_ACCESS_URL_: "you can find this url in the My Account/Control Center API/Access URL tab of your account, and add /v1.0/jsonrpc/push to it, in most cases it is ***https://cloud.gravityzone.bitdefender.com/api/v1.0/jsonrpc/push***"
* _BDGZ_URL_: "url of the server or log collector that will receive the Bitdefender logs. Example: `https://example.com:8000`"

### Execute the Connector

If you are going to run the connector with Docker, you should use this image `docker pull ghcr.io/atlasinsidecorp/bdgz:prod` to build a Docker container with the environment variables defined as explained in the previous step. 

If you are going to run the connector using the executable:
* Define environment variables in an .env file, as described in previous steps
* Build your executable using go:
```
go build
```
* Then run the program passing the run parameter and the connector will be listening on the designated port
```
.\gz_utmstackconn run
```

### Bitdefender configuration

Then, to send your configuration to the Bitdefender Push API, open another cmd console and run the program passing it the sentConfig parameter.
```
.\gz_utmstackconn sentConfig
```

If you want to check your Bitdefender Push API configuration, you should run its executable passing it the getConfig parameter.
```
.\gz_utmstackconn getConfig
```

## Tests

_Then, to send a test log to Bitdefender's Push API, you must run its executable passing it the logTest parameter._
```
.\gz_utmstackconn logTest
```

_At this time, the test logs should be arriving at your server or log collector_

## License
_This project is under AGPL-3.0 license_
 