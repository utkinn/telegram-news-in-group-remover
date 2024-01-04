package db

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type removal struct {
	Message tgbotapi.Message
}

var removalsDb = database[removal]{
	filename: "removals.json",
}

func init() {
	removalsDb.load()
}

func RecordMessageRemoval(message *tgbotapi.Message) {
	removalsDb.add(removal{Message: *message})
}
