package db

import tgbotapi "github.com/utkinn/telegram-bot-api/v5"

type nameReplacement struct {
	Username, NameReplacement string
}

var nameReplacementsDb = database[nameReplacement]{
	filename: "name-replacements.json",
}

func GetNameForUser(user *tgbotapi.User) string {
	for _, repl := range nameReplacementsDb.data {
		if repl.Username == user.UserName {
			return repl.NameReplacement
		}
	}
	return user.FirstName
}
