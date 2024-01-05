package db

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type RemovedMessageDB struct{ database[removal] }

type removal struct {
	Message tgbotapi.Message
}

var removalsDB = RemovedMessageDB{database[removal]{
	filename: "removals.json",
}}

func GetRemovedMessageDB() *RemovedMessageDB {
	return &removalsDB
}

func init() {
	removalsDB.load()
}

func (db *RemovedMessageDB) Add(message *tgbotapi.Message) {
	db.add(removal{Message: *message})
}
