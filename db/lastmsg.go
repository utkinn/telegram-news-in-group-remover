package db

import (
	"log"
	"os"
	"strconv"
)

const lastGroupMessageChatIdFile = "last-group-message-chat-id"

func SetLastMessageChatId(chatId int64) {
	if err := os.WriteFile(lastGroupMessageChatIdFile, []byte(strconv.FormatInt(chatId, 10)), 0644); err != nil {
		log.Printf("Failed to save last group message chat ID: %v\n", err)
	}
}

const LastGroupMessageChatIdNotSet = -1

func LastMessageChatId() int64 {
	idBytes, err := os.ReadFile(lastGroupMessageChatIdFile)
	if err != nil {
		if os.IsNotExist(err) {
			return LastGroupMessageChatIdNotSet
		}
		log.Printf("Failed to load last group message chat ID: %v\n", err)
	}

	idString := string(idBytes)
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		log.Printf("Failed to parse \"%s\" as int64\n", idString)
	}

	return id
}
