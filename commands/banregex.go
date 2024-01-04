package commands

import (
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func init() {
	registerCommand(
		newCommand("banregex", "Пристальный присмотр - запретить сообщения, совпадающие по регулярному выражению", func(ctx helpers.ResponseContext) {
			regex := ctx.Message.CommandArguments()
			if len(regex) == 0 {
				ctx.SendSilentMarkdownFmt("Ты забыл ввести регулярное выражение.\n[Тут](https://golang-blog.blogspot.com/2020/03/regexp-golang.html) можно узнать, что это такое. А [тут](https://regex101.com) можно его потестировать.")
				return
			}

			err := db.BanRegex(regex)
			if err != nil {
				ctx.SendSilentFmt("Не удалось добавить регулярное выражение: %v", err)
				return
			}
			ctx.SendSilentMarkdownFmt("Сообщения от пользователей в списке _пристального присмотра_, совпадающие с этим регулярным выражением, будут удалены.")
		}),
	)
}
