package db

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type NameReplacementDB struct{ database[nameReplacement] }

type nameReplacement struct {
	Username, NameReplacement string
}

var nameReplacementsDb = NameReplacementDB{database[nameReplacement]{
	filename: "name-replacements.json",
}}

func init() {
	nameReplacementsDb.load()
}

func GetNameReplacementDB() *NameReplacementDB {
	return &nameReplacementsDb
}

func (db *NameReplacementDB) GetNameForUser(user *tgbotapi.User) string {
	for _, repl := range db.data {
		if repl.Username == user.UserName {
			return repl.NameReplacement
		}
	}
	return user.FirstName
}
