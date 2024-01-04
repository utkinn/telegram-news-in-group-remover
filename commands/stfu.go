package commands

import (
	"strings"

	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func init() {
	registerCommand(
		newCommand(
			"stfu",
			"Отправляет пользователя с указанным ником на принудительный отдых, запрещая писать ему в беседу на некоторое время",
			func(ctx helpers.ResponseContext) {
				userName := ctx.Message.CommandArguments()
				if len(userName) == 0 || strings.Contains(userName, " ") {
					ctx.SendSilentMarkdownFmt("_Нужен один аргумент — ник пользователя._")
					return
				}

				db.MuteUser(userName)
				ctx.SendSilentMarkdownFmt("Пользователь с ником %s отправлен на принудительный отдых. Если захочется скостить срок — используй /unstfu.", userName)
			},
		),
	)
}
