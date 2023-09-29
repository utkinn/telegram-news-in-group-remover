package db

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const stickersDatabaseFile = "stickers.json"

var stickerFileIds []string

func loadStickersDb() {
	content, err := os.ReadFile(stickersDatabaseFile)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Fatalf("Failed to read %s: %s", stickersDatabaseFile, err.Error())
	}

	if err = json.Unmarshal(content, &stickerFileIds); err != nil {
		log.Fatalf("Failed to unmarshal the contents of %s: %s", stickersDatabaseFile, err.Error())
	}
}

func GetRandomMockStickerFileId() tgbotapi.FileID {
	return tgbotapi.FileID(stickerFileIds[rand.Intn(len(stickerFileIds))])
}
