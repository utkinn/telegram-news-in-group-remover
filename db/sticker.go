package db

import (
	"math/rand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StickerDB struct{ database[string] }

var stickersDb = StickerDB{database[string]{
	filename: "stickers.json",
}}

func init() {
	stickersDb.load()
}

func GetStickerDB() *StickerDB {
	return &stickersDb
}

func (db *StickerDB) GetRandomMockStickerFileId() tgbotapi.FileID {
	return tgbotapi.FileID(stickersDb.data[rand.Intn(len(stickersDb.data))])
}
