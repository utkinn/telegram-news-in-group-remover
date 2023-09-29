package db

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type nameReplacement struct {
	Username, NameReplacement string
}

var nameReplacementsDb = database[nameReplacement]{
	filename: "name-replacement.json",
}

func GetNameForUser(user *tgbotapi.User) string {
	for _, repl := range nameReplacementsDb.data {
		if repl.Username == user.UserName {
			return repl.NameReplacement
		}
	}
	return user.UserName
}