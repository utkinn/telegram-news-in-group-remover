package main

import (
	"time"

	tgbotapi "github.com/utkinn/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

type mock struct {
	messages []*tgbotapi.Message
	time     time.Time
}

const mockCleanupDelay = 2 * time.Minute

func mockCleaner(bot *tgbotapi.BotAPI) {
	for m := range mockCleanupQueue {
		time.Sleep(mockCleanupDelay - (time.Since(m.time)))
		for _, msg := range m.messages {
			helpers.Send(bot, tgbotapi.NewDeleteMessage(msg.Chat.ID, msg.MessageID))
		}
	}
}
