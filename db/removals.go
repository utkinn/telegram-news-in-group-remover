package db

import tgbotapi "github.com/utkinn/telegram-bot-api/v5"

type removal struct {
	Message tgbotapi.Message
}

var removalsDb = database[removal]{
	filename: "removals.json",
}

func RecordMessageRemoval(message *tgbotapi.Message) {
	removalsDb.add(removal{Message: *message})
}
