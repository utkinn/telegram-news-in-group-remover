package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func Clear(ctx helpers.ResponseContext) {
	db.ClearBannedChannels()
	response := tgbotapi.NewMessage(ctx.Message.Chat.ID, "_Список забаненных каналов очищен._")
	response.ParseMode = "markdown"
	helpers.Send(ctx.Bot, response)
}
