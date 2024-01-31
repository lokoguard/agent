package main

import (
	"bytes"
	"errors"
	"net/http"
	"os"
)

func ServerEndpoint() string {
	return "http://localhost:3000"
}

func AuthToken() string {
	auth_token := os.Getenv("LOKOGUARD_CLIENT_TOKEN")
	if auth_token == "" {
		panic("LOKOGUARD_CLIENT_TOKEN not set")
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
	// Send the request
	client := &http.Client{}
	return client.Do(request)
}
