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