package commands

import "github.com/utkinn/telegram-news-in-group-remover/helpers"

func Start(ctx helpers.ResponseContext) {
	ctx.SendSilentMarkdownFmt(
		"Этот бот удаляет сообщения, пересланные из забаненных вами каналов.\n\n" +
			"Для того, чтобы забанить канал, перешлите из него сообщение сюда.\n" +
			"Чтобы очистить список забаненных каналов, выполните /clear.\n" +
			"Посмотреть список забаненных каналов — /list.",
	)
}
