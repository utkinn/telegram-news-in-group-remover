package commands

import (
	"strings"

	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func init() {
	registerCommand(
		newCommand("scrutiny", "<ник> - Начать пристальное внимание за пользователем", func(ctx helpers.ResponseContext) {
			userName := ctx.Message.CommandArguments()
			if len(userName) == 0 || strings.Contains(userName, " ") {
				ctx.SendSilentMarkdownFmt("_Нужен один аргумент — ник пользователя._")
				return
			}
			db.GetScrutinyDB().Add(userName)
			ctx.SendSilentMarkdownFmt("*%s* теперь под _пристальным присмотром_. Вытащить его оттуда можно командой /unscrutiny.", userName)
		}),
	)
}
