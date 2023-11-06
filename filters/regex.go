package filters

import (
	"errors"
	"github.com/dlclark/regexp2"
	tgbotapi "github.com/utkinn/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"log"
)

type regexFilter struct{}

func (f *regexFilter) IsMessageAllowed(message *tgbotapi.Message) bool {
	regexes := db.GetRegexes()
	for _, regex := range regexes {
		caseInsensitiveRegex, err := regexp2.Compile(regex, regexp2.IgnoreCase)
		if err != nil {
			log.Printf("Failed to compile regex %s: %v", caseInsensitiveRegex, err)
			return true
		}
		textMatch, textErr := caseInsensitiveRegex.MatchString(message.Text)
		captionMatch, captionErr := caseInsensitiveRegex.MatchString(message.Caption)
		if captionMatch || textMatch {
			return false
		}
		if err := errors.Join(textErr, captionErr); err != nil {
			log.Printf("regexp %s MatchString failed: %v\nText to match against was: \"%s\"\nCaption to match against was: \"%s\"\n",
				regex, err, message.Text, message.Caption)
		}
	}

	return true
}

func (f *regexFilter) ScrutinyModeOnly() bool {
	return true
}

func (f *regexFilter) ShouldSuppressMock() bool {
	return false
}

func (f *regexFilter) Description() Description {
	return Description{
		ID:   "regex",
		Name: "Регулярные выражения",
		Desc: "Блокирует сообщения, текст которых совпадает с как минимум одным регулярным выражением из запретного списка.",
	}
}
