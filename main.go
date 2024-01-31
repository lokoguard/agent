package main

func main() {
	// Start the syslog server
	StartSyslogServer()
	// Start the resource stats logger
	StartResourceStatsLogger()
	// Start the script executor
	StartScriptExecutorService()
	
	// wait lifetime
	select {}
}
