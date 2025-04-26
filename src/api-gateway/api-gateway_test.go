package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestMainPage(t *testing.T) {
	response, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Errorf("Error with requesting: %v", err)
	}
	defer response.Body.Close()
	if response.Status != "200 OK" {
		t.Errorf("Bad status: \"%v\"", response.Status)
	}
}

func TestLoggingIn(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/auth/sign-in/", bytes.NewBufferString("{username:\"Coolman\",password:\"1234\"}"))
	if err != nil {
		t.Error(err)
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Errorf("Error with requesting: %v", err)
	}
	if response.Status != "200 OK" {
		t.Errorf("Bad status: \"%v\"", response.Status)
	}
}
