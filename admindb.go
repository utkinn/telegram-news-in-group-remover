package main

import (
	"encoding/json"
	"log"
	"os"
)

var adminNicks []string

const adminsDatabaseFile = "admins.json"

func readAdminsDatabase() {
	content, err := os.ReadFile(adminsDatabaseFile)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Fatalf("Failed to read %s: %s", adminsDatabaseFile, err.Error())
	}

	if err = json.Unmarshal(content, &adminNicks); err != nil {
		log.Fatalf("Failed to unmarshal the contents of %s: %s", adminsDatabaseFile, err.Error())
	}
}

func isAdmin(nick string) bool {
	for _, n := range adminNicks {
		if n == nick {
			return true
		}
	}
	return false
}
