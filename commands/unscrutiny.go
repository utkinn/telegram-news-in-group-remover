package commands

import (
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
	"strings"
)

var unscrutinyCommand = newCommand("unscrutiny", "<ник> - Прекратить пристальное внимание за пользователем", func(ctx helpers.ResponseContext) {
	userName := ctx.Message.CommandArguments()
	if len(userName) == 0 || strings.Contains(userName, " ") {
		ctx.SendSilentMarkdownFmt("_Нужен один аргумент — ник пользователя._")
		return
	}
	if db.RemoveFromScrutiny(userName) {
		ctx.SendSilentMarkdownFmt("*%s* выписан из списка _пристального присмотра_.", userName)
	} else {
		ctx.SendSilentMarkdownFmt("*%s* не было в списке _пристального присмотра_, но вычеркнуть я попытался.", userName)
	}
})
