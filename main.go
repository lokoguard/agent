package main

import "os"

func ServerEndpoint() string {
	return "http://localhost:3000"
}

func main() {
	// fetch lokoguard_client_token from environemnt
	auth_token := os.Getenv("LOKOGUARD_CLIENT_TOKEN")
	if auth_token == "" {
		panic("LOKOGUARD_CLIENT_TOKEN not set")
	}
}
