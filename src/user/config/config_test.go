package config

import (
	"log"
	"testing"
)

func TestGetting(t *testing.T) {
	if _, err := GetConfig(); err != nil {
		log.Fatalf("Error with getting config: %v", err)
	}
}
