package commands

import (
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

const confirmation = "да начнется же спам"

func Clear(ctx helpers.ResponseContext) {
	if !db.IsSuperAdmin(ctx.Message.From.UserName) {
		ctx.SendSilentMarkdownFmt("Только владалец может очищать список.")
		return
	}

	if ctx.Message.CommandArguments() != confirmation {
		ctx.SendSilentMarkdownFmt("*Вы в своем уме?*\nОтправьте \"/clear %s\", если вы точно хотите начать хаос.", confirmation)
		return
	}

	db.ClearBannedChannels()
	ctx.SendSilentMarkdownFmt("Ну, как хочешь.\n_Список забаненных каналов очищен._")
}
