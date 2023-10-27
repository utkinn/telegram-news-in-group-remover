package db

import (
	"math/rand"

	tgbotapi "github.com/utkinn/telegram-bot-api/v5"
)

var stickersDb = database[string]{
	filename: "stickers.json",
}

func GetRandomMockStickerFileId() tgbotapi.FileID {
	return tgbotapi.FileID(stickersDb.data[rand.Intn(len(stickersDb.data))])
}
