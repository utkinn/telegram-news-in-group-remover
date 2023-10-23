package db

import (
	"log"
	"os"
	"strconv"
)

const superAdminChatIdFile = "superadmin-chat-id"

func SetSuperAdminChatId(chatId int64) {
	if err := os.WriteFile(superAdminChatIdFile, []byte(strconv.FormatInt(chatId, 10)), 0644); err != nil {
		log.Printf("Failed to save admin chat ID: %v\n", err)
	}
}

func GetSuperAdminChatId() int64 {
	idBytes, err := os.ReadFile(superAdminChatIdFile)
	if err != nil {
		if os.IsNotExist(err) {
			return -1
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
