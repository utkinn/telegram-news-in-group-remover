package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func Unknown(ctx helpers.ResponseContext) {
	text := fmt.Sprintf("*Неизвестная команда: `%s`*", ctx.Message.Command())
	response := tgbotapi.NewMessage(ctx.Message.Chat.ID, text)
	response.ParseMode = "markdown"
	helpers.Send(ctx.Bot, response)
}
