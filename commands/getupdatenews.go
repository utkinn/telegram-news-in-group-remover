package commands

import (
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func init() {
	registerCommand(
		newCommand("getupdatenews", "Подписаться на новости об обновлениях этого бота", func(ctx helpers.ResponseContext) {
			db.GetAnnouncementSubscriptionDB().Subscribe(ctx.Message.Chat.ID, ctx.Message.From.UserName)
			ctx.SendSilentMarkdownFmt("Ты подписан на новости об обновлениях. Отписаться можно командой /noupdatenews.")
		}),
	)
}
