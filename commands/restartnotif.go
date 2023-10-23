package commands

import (
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

var restartNotificationCommand = newHiddenSuperAdminCommand("restartnotif", func(ctx helpers.ResponseContext) {
	db.SetSuperAdminChatId(ctx.Message.Chat.ID)
})