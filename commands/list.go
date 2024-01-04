package commands

import (
	"fmt"

	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func init() {
	const help = "\n\n_Канал можно забанить, переслав любое сообщение из него мне._"

	registerCommand(
		newCommand("list", "Список забаненных каналов", func(ctx helpers.ResponseContext) {
			ctx.SendSilentFmt(listChannelsToString(db.GetBannedChannels()) + help)
		}),
	)
}

func listChannelsToString(bannedChannels []db.Channel) string {
	if len(bannedChannels) == 0 {
		return "Каналов нет"
	}
	lines := make([]any, len(bannedChannels))
	for i := 0; i < len(bannedChannels); i++ {
		lines[i] = fmt.Sprintf("%d. %s\n", i+1, bannedChannels[i].Title)
	}
	return fmt.Sprint(lines...)
}
