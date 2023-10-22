package commands

import (
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
	"strings"
)

var scrutinyCommand = newCommand("scrutiny", "<ник> - Начать пристальное внимание за пользователем", func(ctx helpers.ResponseContext) {
	userName := ctx.Message.CommandArguments()
	if len(userName) == 0 || strings.Contains(userName, " ") {
		ctx.SendSilentMarkdownFmt("_Нужен один аргумент — ник пользователя._")
		return
	}
	db.AddToScrutiny(userName)
	ctx.SendSilentMarkdownFmt("%s теперь под *пристальным присмотром*.", userName)
})
