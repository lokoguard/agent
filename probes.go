package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/lokoguard/agent/resource_monitoring"
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
		jsonData, err := json.Marshal(x)
		if err == nil {
			resp, err := POSTRequest("/api/agent/log", jsonData)
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

func StartResourceStatsLogger() {
	var logger = log.New(os.Stdout, "Resource Stats Logger : ", 0)
	go func() {
		for {
			stats, err := resource_monitoring.FetchResourceStats()
			if err != nil {
				logger.Println(err)
			} else {
				jsonData, err := stats.JSON()
				if err == nil {
					resp, err := POSTRequest("api/agent/monitor/resource", jsonData)
					if err != nil {
						logger.Println(err)
					}
					if resp.StatusCode != 200 {
						logger.Println("Error: ", resp.Status)
						buf := new(bytes.Buffer)
						buf.ReadFrom(resp.Body)
						logger.Println(buf.String())
					}
				} else {
					logger.Println(err)
				}
			}
			// sleep for 10 seconds
			time.Sleep(10 * time.Second)
		}
	}()
	logger.Println("Resource Stats Logger started")
}
