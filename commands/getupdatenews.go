package commands

import (
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

var getUpdateNewsCommand = newCommand("getupdatenews", "Подписаться на новости об обновлениях этого бота", func(ctx helpers.ResponseContext) {
	db.SubscribeToAnnouncements(ctx.Message.Chat.ID, ctx.Message.From.UserName)
	ctx.SendSilentMarkdownFmt("Ты подписан на новости об обновлениях.")
})
