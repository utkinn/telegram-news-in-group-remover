package main

import (
	"fmt"
	"testing"

	tgbotapi "github.com/utkinn/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

type testTextResponder []testResponseRecord

type testResponseRecord struct {
	parseMode, message string
	silent             bool
}

func (r *testTextResponder) RespondTextf(parseMode string, silent bool, format string, a ...any) *tgbotapi.Message {
	*r = append(*r, testResponseRecord{
		parseMode: parseMode,
		silent:    silent,
		message:   fmt.Sprintf(format, a...),
	})
	return nil
}

func TestRejectsUnknownUsersInBotChat(t *testing.T) {
	db.SetAdminsForTesting()

	var responder testTextResponder
	handleMessageToBot(
		helpers.ResponseContext{Message: &tgbotapi.Message{From: &tgbotapi.User{UserName: "RandomDude"}}},
		&responder,
	)

	if len(responder) != 1 {
		t.Fatalf("Want 1 response, got %v", len(responder))
	}
	if responder[0].message != goAway {
		t.Fatalf("Response message: want %v, got %v", goAway, responder[0].message)
	}
}

func TestSkipMessagesToBotWithoutText(t *testing.T) {
	db.SetAdminsForTesting()

	var responder testTextResponder
	handleMessageToBot(
		helpers.ResponseContext{Message: &tgbotapi.Message{
			From:    &tgbotapi.User{UserName: "Admin"},
			Text:    "",
			Caption: "",
		}},
		&responder,
	)

	if len(responder) != 0 {
		t.Fatalf("Want 0 responses, got %v", len(responder))
	}
}
