package commands

import (
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

var banRegexCommand = newCommand("banregex", "Пристальный присмотр - запретить сообщения, совпадающие по регулярному выражению", func(ctx helpers.ResponseContext) {
	regex := ctx.Message.CommandArguments()
	if len(regex) == 0 {
		ctx.SendSilentMarkdownFmt("Ты забыл ввести регулярное выражение.\nСправка [тут](https://golang-blog.blogspot.com/2020/03/regexp-golang.html). Подебажить выражение можешь [тут](https://regex101.com).")
		return
	}

	err := db.BanRegex(regex)
	if err != nil {
		ctx.SendSilentFmt("Не удалось добавить регулярное выражение: %v", err)
		return
	}
	ctx.SendSilentMarkdownFmt("Сообщения от пользователей в списке _пристального присмотра_, совпадающие с этим регулярным выражением, будут удалены.")
})
