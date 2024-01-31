package syslogserver

import (
	"fmt"
	"net"
)

func Start(syslogMessageProcessor ResultCallbackType) {
	// TCP Server
	tcpListener, err := net.Listen("tcp", ":601")
	if err != nil {
		fmt.Println("Error creating TCP listener:", err)
		return
	}
	defer tcpListener.Close()
	go func() {
		for {
			conn, err := tcpListener.Accept()
			if err != nil {
				fmt.Println("Error accepting TCP connection:", err)
				return
			}
			go handleTCP(conn, syslogMessageProcessor)
		}
	}()

	// UDP Server
	udpAddr, err := net.ResolveUDPAddr("udp", ":514")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	udpListener, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error creating UDP listener:", err)
		return
	}
	defer udpListener.Close()
	handleUDP(udpListener, syslogMessageProcessor)

	// Keep the main goroutine running
	select {}
}

func handleTCP(conn net.Conn, syslogMessageProcessor ResultCallbackType) {
	defer conn.Close()
	buffer := make([]byte, 16384)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading TCP:", err)
			return
		}
		message := string(buffer[:n])
		processIncomingSyslogMessage(message, syslogMessageProcessor)
	}
}

func handleUDP(conn *net.UDPConn, syslogMessageProcessor ResultCallbackType) {
	buffer := make([]byte, 4096)
	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading UDP:", err)
			return
		}
		message := string(buffer[:n])
		processIncomingSyslogMessage(message, syslogMessageProcessor)
	}
}

func processIncomingSyslogMessage(message string, syslogMessageProcessor ResultCallbackType) {
	syslogMessage, err := FormatSyslogMessage(message)
	if err != nil {
		fmt.Println("Error parsing syslog message:", err)
		return
	}
	go syslogMessageProcessor(&syslogMessage, nil)
}