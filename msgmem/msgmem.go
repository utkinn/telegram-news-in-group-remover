package msgmem

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Add(m *tgbotapi.Message) {
	messageMemory[ringBufferIndex] = m
	ringBufferIndex = (ringBufferIndex + 1) % len(messageMemory)
}

func Get() []*tgbotapi.Message {
	result := []*tgbotapi.Message{}

	for i := 0; i < len(messageMemory); i++ {
		if messageMemory[i] != nil {
			result = append(result, messageMemory[i])
		}
	}

	return result
}

var (
	messageMemory   [100]*tgbotapi.Message
	ringBufferIndex int
)
