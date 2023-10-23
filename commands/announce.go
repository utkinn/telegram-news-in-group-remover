package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

var announceCommand = newSuperAdminCommand("announce", "Анонсировать обновление", func(ctx helpers.ResponseContext) {
	for _, chatId := range db.GetChatIdsOfAdminsSubscribedToAnnouncements() {
		msg := tgbotapi.NewMessage(chatId, ctx.Message.CommandArguments())
		msg.ParseMode = "markdown"
		msg.DisableNotification = true
		helpers.Send(ctx.Bot, msg)
	}
})
