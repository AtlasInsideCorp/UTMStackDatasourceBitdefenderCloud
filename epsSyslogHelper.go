package main

import (
	"fmt"
	"log"
	"time"

	syslog "github.com/RackSec/srslog"
)

type EpsSyslogHelper struct {
	clientSyslog *syslog.Writer
}

func (g *EpsSyslogHelper) init(config JsonConf) {
	clientSyslog, err := syslog.Dial(config.Protocol, config.Target+":"+config.SyslogPort,
		syslog.LOG_SYSLOG, "bitdefender")
	if err != nil {
		log.Fatalf("Failed create syslog client: %s\n", err)
	}

	g.clientSyslog = clientSyslog
}

func (g *EpsSyslogHelper) sentToSyslog(events []string) {
	for i := 0; i < len(events); i++ {
		syslogMessage := events[i]
		g.clientSyslog.Warning(syslogMessage)
		if getenv("DEBUG") == "true" {
			fmt.Println(time.Now(), ": ", events[i])
		}
	}
}
