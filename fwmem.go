package main

import tgbotapi "github.com/utkinn/telegram-bot-api/v5"

type forwardMemoryItem struct {
	groupChatId, channelId int64
	messageId              int
	from                   tgbotapi.User
}

func newForwardMemoryItem(m *tgbotapi.Message) forwardMemoryItem {
	return forwardMemoryItem{
		groupChatId: m.Chat.ID,
		channelId:   m.ForwardFromChat.ID,
		messageId:   m.MessageID,
		from:        *m.From,
	}
}

var forwardMemory []forwardMemoryItem
