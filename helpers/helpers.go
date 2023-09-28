package helpers

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ResponseContext struct {
	Message *tgbotapi.Message
	Bot     *tgbotapi.BotAPI
}

func Send(bot *tgbotapi.BotAPI, c tgbotapi.Chattable) {
	if _, err := bot.Send(c); err != nil {
		log.Println(err.Error())
	}
}

func (ctx ResponseContext) SendSilentMarkdownFmt(format string, args ...any) {
	msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, fmt.Sprintf(format, args...))
	msg.ParseMode = "markdown"
	msg.DisableNotification = true
	Send(ctx.Bot, msg)
}
