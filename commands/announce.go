package commands

import (
	tgbotapi "github.com/utkinn/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func init() {
	registerCommand(
		newSuperAdminCommand("announce", "Анонсировать обновление", func(ctx helpers.ResponseContext) {
			for _, chatId := range db.GetChatIdsOfAdminsSubscribedToAnnouncements() {
				text := ctx.Message.CommandArguments()
				msg := tgbotapi.NewMessage(chatId, text)
				msg.DisableNotification = true
				copyMarkupFromTextCmdArg(*ctx.Message, &msg, len(text))
				helpers.Send(ctx.Bot, msg)
			}
		}),
	)
}
