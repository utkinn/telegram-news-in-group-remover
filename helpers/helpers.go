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

func Send(bot *tgbotapi.BotAPI, chattable tgbotapi.Chattable) *tgbotapi.Message {
	var (
		msg tgbotapi.Message
		err error
	)

	if msg, err = bot.Send(chattable); err != nil {
		log.Println(err.Error())
		return nil
	}

	return &msg
}

func (ctx ResponseContext) SendSilentFmt(format string, args ...any) *tgbotapi.Message {
	msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, fmt.Sprintf(format, args...))
	msg.DisableNotification = true

	return Send(ctx.Bot, msg)
}

func (ctx ResponseContext) SendSilentMarkdownFmt(format string, args ...any) *tgbotapi.Message {
	msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, fmt.Sprintf(format, args...))
	msg.ParseMode = tgbotapi.ModeMarkdown
	msg.DisableNotification = true

	return Send(ctx.Bot, msg)
}
