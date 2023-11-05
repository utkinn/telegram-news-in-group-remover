package commands

import (
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
	"strings"
)

var unmuteCommand = newCommand(
	"unstfu",
	"Возвращает пользователя с указанным ником из принудительного отдыха",
	func(ctx helpers.ResponseContext) {
		userName := ctx.Message.CommandArguments()
		if len(userName) == 0 || strings.Contains(userName, " ") {
			ctx.SendSilentMarkdownFmt("_Нужен один аргумент — ник пользователя._")
			return
		}

		db.UnmuteUser(userName)
		ctx.SendSilentMarkdownFmt("Пользователь с ником %s возвращен с принудительного отдыха.", userName)
	},
)
