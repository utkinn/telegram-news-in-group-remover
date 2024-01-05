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

var stickersDB = StickerDB{
	database[string]{
		filename: "stickers.json",
	},
	realRandom{},
}

func init() {
	stickersDB.load()
}

func GetStickerDB() *StickerDB {
	return &stickersDB
}

func (db *StickerDB) GetRandomMockStickerFileID() tgbotapi.FileID {
	return tgbotapi.FileID(db.data[db.random.Intn(len(db.data))])
}
