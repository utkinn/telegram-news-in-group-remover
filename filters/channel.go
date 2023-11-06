package filters

import (
	tgbotapi "github.com/utkinn/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/db"
)

type channelFilter struct{}

func (f *channelFilter) IsMessageAllowed(message *tgbotapi.Message) bool {
	if message.ForwardFromChat == nil {
		return true
	}
	return !db.IsChannelIdBanned(message.ForwardFromChat.ID)
}

func (f *channelFilter) ScrutinyModeOnly() bool {
	return false
}

func (f *channelFilter) ShouldSuppressMock() bool {
	return false
}

func (f *channelFilter) Description() Description {
	return Description{
		ID:   "channels",
		Name: "Пересылки из забаненных каналов",
		Desc: "Блокирует все пересланные сообщения из заблокированных каналов (/list).",
	}
}
