package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lokoguard/agent/file_monitoring"
	"github.com/lokoguard/agent/resource_monitoring"
	"github.com/lokoguard/agent/script_executor"
	syslogserver "github.com/lokoguard/agent/syslog_server"
)

// All the probes should be non-blocking
// So use goroutines inside the probes implementation

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

func StartScriptExecutorService() {
	var logger = log.New(os.Stdout, "Script Executor Service : ", 0)
	go func() {
		lastRequestsCount := 0
		for {
			// fetch script definitions
			scriptDefinitions, err := fetchScriptDefinitions()
			if err != nil {
				logger.Println(err)
			} else {
				lastRequestsCount = len(scriptDefinitions)
				for _, scriptDefinition := range scriptDefinitions {
					scriptDefinition.RunWithCallback(func(result *script_executor.ScriptResult, err error) {
						if err != nil {
							logger.Println(err)
						} else {
							jsonData, err := result.JSON()
							if err == nil {
								resp, err := POSTRequest("/api/agent/executor/submit", jsonData)
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
					})
				}
			}

			if lastRequestsCount == 0 {
				// sleep for 2 seconds
				time.Sleep(2 * time.Second)
			}
		}
	}()
	logger.Println("Script Executor Service started")
}

func StartFileMonitoringService() {
	var logger = log.New(os.Stdout, "File Monitoring Service : ", 0)
	// Start the file monitoring service
	file_monitor := file_monitoring.NewMonitor(func(event *file_monitoring.UpdateEvent) {
		jsonData, err := event.JSON()
		if err == nil {
			resp, err := POSTRequest("/api/agent/monitor/file", jsonData)
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
	})

	// Start listener for required file tracking infos
	go func() {
		for {
			// fetch file list
			resp, err := GETRequest("/api/agent/monitor/file")
			if err != nil {
				logger.Println(err)
			} else {
				if resp.StatusCode != 200 {
					logger.Println("Error: ", resp.Status)
					buf := new(bytes.Buffer)
					buf.ReadFrom(resp.Body)
					logger.Println(buf.String())
				} else {
					buf := new(bytes.Buffer)
					buf.ReadFrom(resp.Body)
					var fileNames []string
					err = json.Unmarshal(buf.Bytes(), &fileNames)
					if err != nil {
						logger.Println(err)
					} else {
						file_monitor.UpdateFileList(fileNames)
					}
				}
			}
			// sleep for 10 seconds
			time.Sleep(10 * time.Second)
		}
	}()

	logger.Println("File Monitoring Service started")
}

func StartPingService() {
	var logger = log.New(os.Stdout, "Ping Service : ", 0)
	go func() {
		for {
			_, err := GETRequest("/api/agent/ping")
			if err != nil {
				logger.Println(err)
			}
			// sleep for 10 seconds
			time.Sleep(10 * time.Second)
		}
	}()
	logger.Println("Ping Service started")
}

// Helper functions
func fetchScriptDefinitions() ([]script_executor.ScriptDefinition, error) {
	resp, err := GETRequest("/api/agent/executor")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	var scriptDefinitions []script_executor.ScriptDefinition
	err = json.Unmarshal(buf.Bytes(), &scriptDefinitions)
	if err != nil {
		return nil, err
	}
	return scriptDefinitions, nil
}
