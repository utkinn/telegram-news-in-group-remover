package db

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type removal struct {
	Message tgbotapi.Message
	Manual  bool
}

var removalsDb = database[removal]{
	filename: "removals.json",
}

type MsgRemovalType int

const (
	MsgRemovalManual MsgRemovalType = iota
	MsgRemovalAuto
)

func RecordMessageRemoval(message *tgbotapi.Message, removalType MsgRemovalType) {
	removalsDb.add(removal{Message: *message, Manual: removalType == MsgRemovalManual})
}
