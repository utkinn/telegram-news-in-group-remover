package commands

import (
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func init() {
	registerCommand(
		newCommand(
			"noupdatenews",
			"Отписаться от новостей об обновлениях этого бота",
			func(ctx helpers.ResponseContext) {
				db.GetAnnouncementSubscriptionDB().Unsubscribe(ctx.Message.Chat.ID)
				ctx.SendSilentMarkdownFmt("Ты теперь _не_ подписан на новости об обновлениях.")
			},
		),
	)
}
