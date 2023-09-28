package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func Start(ctx helpers.ResponseContext) {
	helpers.Send(
		ctx.Bot,
		tgbotapi.NewMessage(
			ctx.Message.Chat.ID,
			"Этот бот удаляет сообщения, пересланные из забаненных вами каналов.\n\n"+
				"Для того, чтобы забанить канал, перешлите из него сообщение сюда.\n"+
				"Чтобы очистить список забаненных каналов, выполните /clear.\n"+
				"Посмотреть список забаненных каналов — /list.",
		),
	)
}
