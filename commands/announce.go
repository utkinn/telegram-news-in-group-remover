package commands

import (
	tgbotapi "github.com/utkinn/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

var announceCommand = newSuperAdminCommand("announce", "Анонсировать обновление", func(ctx helpers.ResponseContext) {
	for _, chatId := range db.GetChatIdsOfAdminsSubscribedToAnnouncements() {
		msg := tgbotapi.NewMessage(chatId, ctx.Message.CommandArguments())
		msg.DisableNotification = true

		msg.Entities = make([]tgbotapi.MessageEntity, 0, len(ctx.Message.Entities)-1) // -1 for bot_command entity
		for _, ent := range ctx.Message.Entities {
			if ent.Type != "bot_command" {
				ent.Offset -= len(ctx.Message.Command()) + 2 // 2 = "/" and a space after the command
				msg.Entities = append(msg.Entities, ent)
			}
		}

		helpers.Send(ctx.Bot, msg)
	}
})
