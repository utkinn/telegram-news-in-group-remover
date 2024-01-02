package filters

import (
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func init() {
	registerFilter(&storiesFilter{})
}

type storiesFilter struct{}

func (s *storiesFilter) IsMessageAllowed(ctx helpers.ResponseContext) bool {
	return ctx.Message.Story == nil
}

func (s *storiesFilter) ScrutinyModeOnly() bool {
	return true
}

func (s *storiesFilter) ShouldSuppressMock() bool {
	return false
}

func (s *storiesFilter) Description() Description {
	return Description{
		ID:   "stories",
		Name: "Истории",
		Desc: "Запрещает пересылку всех историй. Фильтрация историй по каналам невозможна из-за текущих ограничений API Telegram.",
	}
}
