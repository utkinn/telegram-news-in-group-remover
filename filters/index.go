package filters

import tgbotapi "github.com/utkinn/telegram-bot-api/v5"

type filter interface {
	IsMessageAllowed(message *tgbotapi.Message) bool
	ScrutinyModeOnly() bool
}

var filters = []filter{
	&channelFilter{},
	&regexFilter{},
	&storiesFilter{},
}

func IsMessageAllowed(message *tgbotapi.Message) bool {
	for _, f := range filters {
		if !f.IsMessageAllowed(message) {
			return false
		}
	}
	return true
}
