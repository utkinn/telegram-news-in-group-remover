package commands

import "github.com/utkinn/telegram-news-in-group-remover/helpers"

func Unknown(ctx helpers.ResponseContext) {
	ctx.SendSilentMarkdownFmt("*Неизвестная команда: `%s`*", ctx.Message.Command())
}
