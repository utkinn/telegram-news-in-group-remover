package db

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TestGetNameForUserWorksWithReplacement(t *testing.T) {
	db := NameReplacementDB{
		database[nameReplacement]{
			data: []nameReplacement{
				{Username: "Alex123", NameReplacement: "Funny guy"},
			},
		},
	}

	userWithReplacement := tgbotapi.User{UserName: "Alex123", FirstName: "Alex"}
	name := db.GetNameForUser(&userWithReplacement)
	if name != "Funny guy" {
		t.Fatalf("Name replacement failed, got %v instead", name)
	}

	userWithoutReplacement := tgbotapi.User{UserName: "foobar", FirstName: "John"}
	name = db.GetNameForUser(&userWithoutReplacement)
	if name != "John" {
		t.Fatalf("Name replacement failed, got %v instead", name)
	}
}
