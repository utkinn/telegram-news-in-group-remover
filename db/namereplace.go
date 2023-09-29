package db

import (
	"encoding/json"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type nameReplacement struct {
	Username, NameReplacement string
}

var nameReplacements []nameReplacement

const nameReplacementFile = "name-replacements.json"

func loadNameReplacementsDb() {
	content, err := os.ReadFile(nameReplacementFile)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Fatalf("Failed to read %s: %s", nameReplacementFile, err.Error())
	}

	if err = json.Unmarshal(content, &nameReplacements); err != nil {
		log.Fatalf("Failed to unmarshal the contents of %s: %s", nameReplacementFile, err.Error())
	}
}

func GetNameForUser(user *tgbotapi.User) string {
	for _, repl := range nameReplacements {
		if repl.Username == user.UserName {
			return repl.NameReplacement
		}
	}
	return user.UserName
}
