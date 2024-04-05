package main

func main() {
	// Start the syslog server
	StartSyslogServer()
	// Start the resource stats logger
	StartResourceStatsLogger()
	// Start the script executor
	StartScriptExecutorService()
	// Start the file monitoring service
	StartFileMonitoringService()
	// Start the ping service
	StartPingService()
	// wait lifetime
	select {}
}
