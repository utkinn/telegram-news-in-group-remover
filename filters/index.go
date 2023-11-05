package filters

import tgbotapi "github.com/utkinn/telegram-bot-api/v5"

type filter interface {
	IsMessageAllowed(message *tgbotapi.Message) bool
	ScrutinyModeOnly() bool
	ShouldSuppressMock() bool
}

var filters []filter

func InitFilters(bot *tgbotapi.BotAPI) {
	filters = []filter{
		&channelFilter{},
		&regexFilter{},
		&storiesFilter{},
		&muteFilter{bot: bot},
	}
}

func IsMessageAllowed(message *tgbotapi.Message) (allowed, suppressMock bool) {
	for _, f := range filters {
		if !f.IsMessageAllowed(message) {
			return false, f.ShouldSuppressMock()
		}
	}
	return true, false
}
