package filters

import (
	_ "embed"
	"fmt"
	tgbotapi "github.com/utkinn/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

type muteFilter struct {
	bot *tgbotapi.BotAPI
}

//go:embed chill.png
var chillPng []byte

func (m *muteFilter) IsMessageAllowed(message *tgbotapi.Message) bool {
	senderUserName := message.From.UserName
	muted, announced := db.IsUserMuted(senderUserName)
	if muted && !announced {
		muteAnnouncement := tgbotapi.NewPhoto(message.Chat.ID, tgbotapi.FileBytes{Name: "chill.png", Bytes: chillPng})
		muteAnnouncement.Caption = fmt.Sprintf("@%v, иди паспи, не скажу насколько", senderUserName)
		helpers.Send(m.bot, muteAnnouncement)
		db.MarkMuteAnnounced(senderUserName)
	}
	return !muted
}

func (m *muteFilter) ScrutinyModeOnly() bool {
	return false
}

func (m *muteFilter) ShouldSuppressMock() bool {
	return true
}

func (m *muteFilter) Description() Description {
	return Description{
		ID:   "mute",
		Name: "STFU",
		Desc: "Удаляет все сообщения от пользователей, которые были отправлены в мут командой /stfu.",
	}
}