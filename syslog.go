package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	syslogserver "github.com/lokoguard/agent/syslog_server"
)

func StartSyslogServer() {
	var logger = log.New(os.Stdout, "Syslog Server : ", 0)

	syslogMessageHandler := func(x *syslogserver.SyslogMessage, err error) {
		if err != nil {
			logger.Println(err)
			return
		}
		// json marshal
		jsonText, err := json.Marshal(x)
		if err == nil {
			resp, err := POSTRequest("/api/agent/log", jsonText)
			if err != nil {
				logger.Println(err)
			}
			if resp.StatusCode != 200 {
				logger.Println("Error: ", resp.Status)
				// print body
				buf := new(bytes.Buffer)
				buf.ReadFrom(resp.Body)
				logger.Println(buf.String())
			}
		} else {
			logger.Println(err)
		}
	}

	// Start the syslog server
	go syslogserver.Start(syslogMessageHandler)
	logger.Println("Syslog server started")
}
