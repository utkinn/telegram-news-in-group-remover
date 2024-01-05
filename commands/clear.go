package commands

import (
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

const confirmation = "да начнется же спам"

func init() {
	registerCommand(
		newSuperAdminCommand("clear", "Разбанить все каналы", func(ctx helpers.ResponseContext) {
			if ctx.Message.CommandArguments() != confirmation {
				ctx.SendSilentMarkdownFmt(
					"*Вы в своем уме?*\nОтправьте \"/clear %s\", если вы точно хотите начать хаос.",
					confirmation,
				)
				return
			}

			db.GetBannedChannelDB().Clear()
			ctx.SendSilentMarkdownFmt("Ну, как хочешь.\n_Список забаненных каналов очищен._")
		}),
	)
}
