package main

import (
	"encoding/json"
	"fmt"

	syslogserver "github.com/lokoguard/agent/syslog_server"
)

func ServerEndpoint() string {
	return "http://localhost:3000"
}

func main() {
	// // fetch lokoguard_client_token from environemnt
	// auth_token := os.Getenv("LOKOGUARD_CLIENT_TOKEN")
	// if auth_token == "" {
	// 	panic("LOKOGUARD_CLIENT_TOKEN not set")
	// }

	syslogserver.Start(pr)
}

func pr(x *syslogserver.SyslogMessage, err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
	// json marshal
	jsonText, err := json.Marshal(x)
	if err == nil {
		fmt.Println(string(jsonText))
	}
}
