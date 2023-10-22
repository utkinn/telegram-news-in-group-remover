package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/db"
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

	return true
}
