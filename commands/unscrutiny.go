package commands

import (
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
	"strings"
)

var unscrutinyCommand = newCommand("unscrutiny", "<ник> - Прекратить пристальное внимание за пользователем", func(ctx helpers.ResponseContext) {
	args := ctx.Message.CommandArguments()
	if len(args) == 0 || strings.Contains(args, " ") {
		ctx.SendSilentMarkdownFmt("_Нужен один аргумент — ник пользователя._")
		return
	}
	db.RemoveFromScrutiny(args)
})
