package db

import (
	"path"
	"reflect"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TestRemovedMessageDBAdd(t *testing.T) {
	database := RemovedMessageDB{
		database[removal]{
			filename: path.Join(t.TempDir(), "test-removed-messages.json"),
		},
	}

	message := tgbotapi.Message{
		MessageID: 123,
		Text:      "Test message",
	}

	database.Add(&message)

	if len(database.data) != 1 {
		t.Fatalf("Expected 1 message, got %v", len(database.data))
	}

	addedRemoval := database.data[0]
	if !reflect.DeepEqual(addedRemoval.Message, message) {
		t.Fatalf("Expected message %v, got %v", message, addedRemoval.Message)
	}
}
