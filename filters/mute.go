package filters

import (
	_ "embed"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func init() {
	registerFilter(&muteFilter{})
}

type muteFilter struct{}

//go:embed chill.png
var chillPng []byte

func (*muteFilter) IsMessageAllowed(ctx helpers.ResponseContext) bool {
	senderUserName := ctx.Message.From.UserName

	muted, announced := db.GetMuteDB().GetStatusForUser(senderUserName)
	if muted && !announced {
		muteAnnouncement := tgbotapi.NewPhoto(
			ctx.Message.Chat.ID,
			tgbotapi.FileBytes{Name: "chill.png", Bytes: chillPng},
		)
		muteAnnouncement.Caption = fmt.Sprintf("@%v, иди паспи, не скажу насколько", senderUserName)
		helpers.Send(ctx.Bot, muteAnnouncement)
		db.GetMuteDB().MarkMuteAnnounced(senderUserName)
	}

	return !muted
}

func (*muteFilter) ScrutinyModeOnly() bool {
	return false
}

func (*muteFilter) ShouldSuppressMock() bool {
	return true
}

func (*muteFilter) Description() Description {
	return Description{
		ID:   "mute",
		Name: "STFU",
		Desc: "Удаляет все сообщения от пользователей, которые были отправлены в мут командой /stfu.",
	}
}
