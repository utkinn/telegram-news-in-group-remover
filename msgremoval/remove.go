package msgremoval

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func Remove(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	helpers.Send(bot, tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID))
	db.GetRemovedMessageDB().Add(message)
}
