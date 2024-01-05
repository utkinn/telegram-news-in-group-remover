package commands

import (
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func init() {
	registerCommand(
		newCommand(
			"unbanregex",
			"Пристальный присмотр - убрать регулярное выражение из списка запрещенных",
			func(ctx helpers.ResponseContext) {
				regex := ctx.Message.CommandArguments()
				if len(regex) == 0 {
					ctx.SendSilentMarkdownFmt("Ты забыл ввести регулярное выражение.\n" +
						"Справка [тут](https://golang-blog.blogspot.com/2020/03/regexp-golang.html).")
					return
				}

				db.GetBannedRegexDB().Unban(regex)
				ctx.SendSilentMarkdownFmt("Теперь это регулярное выражение не под запретом.")
			},
		),
	)
}
