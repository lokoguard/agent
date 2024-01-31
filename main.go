package main

func main() {
	// Start the syslog server
	StartSyslogServer()

	// wait lifetime
	select {}
}
