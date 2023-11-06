package filters

import (
	tgbotapi "github.com/utkinn/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/db"
)

type Filter interface {
	IsMessageAllowed(message *tgbotapi.Message) bool
	ScrutinyModeOnly() bool
	ShouldSuppressMock() bool
	Description() Description
}

type Description struct {
	ID, Name, Desc string
}

// Use this notice in Filter.Description to indicate that a certain filter is unstable.
const unstableNotice = "\n*Этот фильтр экспериментален и может работать нестабильно, с большим количеством ложных срабатываний. Не забывайте про команду `/filteroff`.*"

var filters []Filter

func Init(bot *tgbotapi.BotAPI) {
	filters = []Filter{
		&channelFilter{},
		&regexFilter{},
		&storiesFilter{},
		&muteFilter{bot: bot},
		&screenshotFilter{bot: bot},
	}
}

func List() []Filter {
	return filters
}

func IsMessageAllowed(message *tgbotapi.Message) (allowed, suppressMock bool) {
	senderIsUnderScrutiny := db.IsUnderScrutiny(message.From.UserName)
	for _, f := range filters {
		if f.ScrutinyModeOnly() && !senderIsUnderScrutiny {
			continue
		}
		if !f.IsMessageAllowed(message) {
			return false, f.ShouldSuppressMock()
		}
	}
	return true, false
}
