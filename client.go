package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func ServerEndpoint() string {
	endpoint := os.Getenv("lokoguard_agent_endpoint")
	if strings.Compare(endpoint, "") == 0 {
		return "http://localhost:3000"
	}
	return endpoint
}

func AuthToken() string {
	auth_token := os.Getenv("lokoguard_agent_token")
	if auth_token == "" {
		panic("lokoguard_agent_token not set")
	}
	return auth_token
}

func GETRequest(route string) (*http.Response, error) {
	if route == "" {
		return nil, errors.New("route cannot be empty")
	}
	if route[0] != '/' {
		route = "/" + route
	}
	request, err := http.NewRequest("GET", ServerEndpoint()+route, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bearer "+AuthToken())
	request.Header.Set("Accept", "application/json")
	// Send the request
	client := &http.Client{}
	return client.Do(request)
}

func POSTRequest(route string, body []byte) (*http.Response, error) {
	if route == "" {
		return nil, errors.New("route cannot be empty")
	}
	if route[0] != '/' {
		route = "/" + route
	}
	reqBody := bytes.NewBuffer(body)
	request, err := http.NewRequest("POST", ServerEndpoint()+route, reqBody)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bearer "+AuthToken())
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Content-Length", fmt.Sprint(len(body)))
	request.Header.Set("Accept", "application/json")
	// Send the request
	client := &http.Client{}
	return client.Do(request)
}
