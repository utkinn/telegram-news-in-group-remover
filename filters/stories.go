package filters

import (
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func init() {
	registerFilter(&storiesFilter{})
}

type storiesFilter struct{}

func (*storiesFilter) IsMessageAllowed(ctx helpers.ResponseContext) bool {
	return ctx.Message.Story == nil
}

func (*storiesFilter) ScrutinyModeOnly() bool {
	return true
}

func (*storiesFilter) ShouldSuppressMock() bool {
	return false
}

func (*storiesFilter) Description() Description {
	return Description{
		ID:   "stories",
		Name: "Истории",
		Desc: "Запрещает пересылку всех историй. " +
			"Фильтрация историй по каналам невозможна из-за текущих ограничений API Telegram.",
	}
}
