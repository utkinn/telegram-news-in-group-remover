package filters

import (
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func init() {
	registerFilter(&channelFilter{})
}

type channelFilter struct{}

func (f *channelFilter) IsMessageAllowed(ctx helpers.ResponseContext) bool {
	if ctx.Message.ForwardFromChat == nil {
		return true
	}
	return !db.GetBannedChannelDB().IsBanned(ctx.Message.ForwardFromChat.ID)
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
