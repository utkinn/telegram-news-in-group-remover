package commands

import (
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
	"strings"
)

var scrutinyCommand = newCommand("scrutiny", "<ник> - Начать пристальное внимание за пользователем", func(ctx helpers.ResponseContext) {
	args := ctx.Message.CommandArguments()
	if strings.Contains(args, " ") {
		ctx.SendSilentMarkdownFmt("_Нужен один аргумент — ник пользователя._")
		return
	}
	db.AddToScrutiny(args)
})
