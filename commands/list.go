package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func List(ctx helpers.ResponseContext) {
	bannedChannels := db.GetBannedChannels()
	text := listChannelsToString(bannedChannels)
	response := tgbotapi.NewMessage(ctx.Message.Chat.ID, text)
	response.ParseMode = "markdown"
	helpers.Send(ctx.Bot, response)
}

func listChannelsToString(bannedChannels []db.Channel) string {
	if len(bannedChannels) == 0 {
		return "_Каналов нет_"
	}
	lines := make([]any, len(bannedChannels))
	for i := 0; i < len(bannedChannels); i++ {
		lines[i] = fmt.Sprintf("%d. %s\n", i+1, bannedChannels[i].Title)
	}
	return fmt.Sprint(lines...)
}
