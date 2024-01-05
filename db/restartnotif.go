package db

import (
	"log"
	"os"
	"strconv"
)

const superAdminChatIDFile = "superadmin-chat-id"

func SetSuperAdminChatID(chatID int64) {
	if err := os.WriteFile(superAdminChatIDFile, []byte(strconv.FormatInt(chatID, 10)), 0600); err != nil {
		log.Printf("Failed to save admin chat ID: %v\n", err)
	}
}

const SuperAdminChatIDNotSet = -1

func GetSuperAdminChatID() int64 {
	idBytes, err := os.ReadFile(superAdminChatIDFile)
	if err != nil {
		if os.IsNotExist(err) {
			return SuperAdminChatIDNotSet
		}

		log.Printf("Failed to load admin chat ID: %v\n", err)
	}

	idString := string(idBytes)

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		log.Printf("Failed to parse \"%s\" as int64\n", idString)
	}

	return id
}
