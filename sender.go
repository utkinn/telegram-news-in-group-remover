package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type textResponder interface {
	RespondTextf(parseMode string, silent bool, format string, a ...any) *tgbotapi.Message
}

type telegramTextResponder struct {
	bot    *tgbotapi.BotAPI
	chatID int64
}

func newTelegramTextResponder(bot *tgbotapi.BotAPI, chatID int64) telegramTextResponder {
	return telegramTextResponder{bot: bot, chatID: chatID}
}

func (r telegramTextResponder) RespondTextf(parseMode string, silent bool, format string, a ...any) *tgbotapi.Message {
	msg := tgbotapi.NewMessage(r.chatID, fmt.Sprintf(format, a...))
	msg.ParseMode = parseMode
	msg.DisableNotification = silent

	var (
		sentResponse tgbotapi.Message
		err          error
	)

	if sentResponse, err = r.bot.Send(msg); err != nil {
		log.Println(err.Error())
		return nil
	}

	return &sentResponse
}
