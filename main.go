package main

func main() {
	// Start the syslog server
	StartSyslogServer()
	// Start the resource stats logger
	StartResourceStatsLogger()

	// wait lifetime
	select {}
}
