package main

import (
	"errors"
	"github.com/dlclark/regexp2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"log"
	"strings"
)

func passesScrutinyFilters(msg *tgbotapi.Message) bool {
	if !db.IsUnderScrutiny(msg.From.UserName) {
		return true
	}

	// Ban suspected copies of news texts
	if strings.Contains(msg.Text, "\n\n") {
		return false
	}

	if matchesBannedRegexes(msg.Text, msg.Caption) {
		return false
	}

	return true
}

func matchesBannedRegexes(text, caption string) bool {
	// TODO: remove messages AND attachments
	regexes := db.GetRegexes()
	for _, regex := range regexes {
		caseInsensitiveRegex, err := regexp2.Compile(regex, regexp2.IgnoreCase)
		if err != nil {
			log.Printf("Failed to compile regex %s: %v", caseInsensitiveRegex, err)
		}
		textMatch, textErr := caseInsensitiveRegex.MatchString(text)
		captionMatch, captionErr := caseInsensitiveRegex.MatchString(caption)
		if captionMatch || textMatch {
			return true
		}
		if err := errors.Join(textErr, captionErr); err != nil {
			log.Printf("regexp %s MatchString failed: %v\nText to match against was: \"%s\"\nCaption to match against was: \"%s\"\n", regex, err, text)
		}
	}

	return false
}
