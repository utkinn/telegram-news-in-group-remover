package filters

import (
	"log"

	"github.com/dlclark/regexp2"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func init() {
	registerFilter(&regexFilter{})
}

type regexFilter struct{}

func (*regexFilter) IsMessageAllowed(ctx helpers.ResponseContext) bool {
	regexes := db.GetBannedRegexDB().Get()

	for _, regex := range regexes {
		compiledRegex, err := regexp2.Compile(regex, regexp2.IgnoreCase)
		if err != nil {
			log.Printf("Failed to compile regex %s: %v", compiledRegex, err)
			return true
		}

		match, err := compiledRegex.MatchString(ctx.Message.Text + ctx.Message.Caption)
		if err != nil {
			log.Printf("regexp %s MatchString failed: %v\nText to match against was: \"%s\"\n"+
				"Caption to match against was: \"%s\"\n",
				regex, err, ctx.Message.Text, ctx.Message.Caption)
		}

		if match {
			return false
		}
	}

	return true
}

func (*regexFilter) ScrutinyModeOnly() bool {
	return true
}

func (*regexFilter) ShouldSuppressMock() bool {
	return false
}

func (*regexFilter) Description() Description {
	return Description{
		ID:   "regex",
		Name: "Регулярные выражения",
		Desc: "Блокирует сообщения, текст которых совпадает " +
			"с как минимум одным регулярным выражением из запретного списка (/listregex). " +
			"Команды: /banregex, /unbanregex",
	}
}
