package db

import (
	"math/rand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type random interface {
	Intn(n int) int
}

type realRandom struct{}

func (realRandom) Intn(n int) int { return rand.Intn(n) }

type StickerDB struct {
	database[string]
	random random
}

var stickersDb = StickerDB{
	database[string]{
		filename: "stickers.json",
	},
	realRandom{},
}

func init() {
	stickersDb.load()
}

func GetStickerDB() *StickerDB {
	return &stickersDb
}

func (db *StickerDB) GetRandomMockStickerFileId() tgbotapi.FileID {
	return tgbotapi.FileID(db.data[db.random.Intn(len(db.data))])
}
