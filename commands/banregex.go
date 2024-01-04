package commands

import (
	"log"

	"github.com/dlclark/regexp2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
	"github.com/utkinn/telegram-news-in-group-remover/msgmem"
	"github.com/utkinn/telegram-news-in-group-remover/msgremoval"
)

func init() {
	registerCommand(
		newCommand("banregex", "Пристальный присмотр - запретить сообщения, совпадающие по регулярному выражению", func(ctx helpers.ResponseContext) {
			regex := ctx.Message.CommandArguments()
			if len(regex) == 0 {
				ctx.SendSilentMarkdownFmt("Ты забыл ввести регулярное выражение.\n" +
					"[Тут](https://golang-blog.blogspot.com/2020/03/regexp-golang.html) можно узнать, что это такое. " +
					"А [тут](https://regex101.com) можно его потестировать.")
				return
			}

			err := db.BanRegex(regex)
			if err != nil {
				ctx.SendSilentFmt("Не удалось добавить регулярное выражение: %v", err)
				return
			}

			removeMatchingMessagesFromMsgmem(regex, ctx.Bot)

			ctx.SendSilentMarkdownFmt("Сообщения от пользователей в списке _пристального присмотра_, совпадающие с этим регулярным выражением, будут удалены.")
		}),
	)
}

func removeMatchingMessagesFromMsgmem(regex string, bot *tgbotapi.BotAPI) {
	compiledRegex, err := regexp2.Compile(regex, regexp2.IgnoreCase)
	if err != nil {
		log.Printf("Failed to compile regex %s, skipping removing matching messages from msgmem: %v", compiledRegex, err)
		return
	}

	userNamesMockedSoFar := map[string]bool{}
	for _, item := range msgmem.Get() {
		match, err := compiledRegex.MatchString(item.Text)
		if err != nil {
			log.Printf("regexp %s MatchString failed: %v\nText to match against was: \"%s\"\n",
				regex, err, item.Text)
			continue
		}
		if match && db.IsUnderScrutiny(item.From.UserName) { // TODO: check using filter
			msgremoval.Remove(bot, item)
			if !userNamesMockedSoFar[item.From.UserName] {
				msgremoval.MockUser(bot, item.Chat.ID, item.From)
				userNamesMockedSoFar[item.From.UserName] = true
			}
		}
	}
}
