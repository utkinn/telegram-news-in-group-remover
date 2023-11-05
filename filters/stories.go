package filters

import (
	tgbotapi "github.com/utkinn/telegram-bot-api/v5"
)

type storiesFilter struct{}

func (s *storiesFilter) IsMessageAllowed(message *tgbotapi.Message) bool {
	return message.Story == nil
}

func (s *storiesFilter) ScrutinyModeOnly() bool {
	return true
}

func (s *storiesFilter) ShouldSuppressMock() bool {
	return false
}
