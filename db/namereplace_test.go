package db

import (
	"os"
	"path"
	"testing"

	tgbotapi "github.com/utkinn/telegram-bot-api/v5"
)

func TestGetNameForUserWorksWithReplacement(t *testing.T) {
	json := `
	[
		{
			"Username": "Alex123",
			"NameReplacement": "Funny guy"
		},
		{
			"Username": "foobar",
			"NameReplacement": "Another one"
		}
	]	
	`
	tempJsonFileName := path.Join(t.TempDir(), "name-replacements.json")
	if err := os.WriteFile(tempJsonFileName, []byte(json), 0444); err != nil {
		t.Fatalf("Failed to write to file: %v", err)
	}
	nameReplacementsDb.filename = tempJsonFileName
	nameReplacementsDb.load()

	user := tgbotapi.User{UserName: "Alex123", FirstName: "Alex"}
	replaced := GetNameForUser(&user)

	if replaced != "Funny guy" {
		t.Fatalf("Name replacement failed, got %v instead", replaced)
	}
}
