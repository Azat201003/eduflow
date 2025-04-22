package main

import (
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
