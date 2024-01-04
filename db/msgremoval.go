package db

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type RemovedMessageDB struct{ database[removal] }

type removal struct {
	Message tgbotapi.Message
}

var removalsDb = RemovedMessageDB{database[removal]{
	filename: "removals.json",
}}

func GetRemovedMessageDB() *RemovedMessageDB {
	return &removalsDb
}

func init() {
	removalsDb.load()
}

func (db *RemovedMessageDB) Add(message *tgbotapi.Message) {
	db.add(removal{Message: *message})
}
