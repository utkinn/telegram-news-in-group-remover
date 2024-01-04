package main

import (
	"fmt"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
	var responder testTextResponder
	handleMessageToBot(
		newResponseContextStub("Chiga", "/clear"),
		&responder,
		func(ctx helpers.ResponseContext) { t.Fatal("commands.Execute was not expected to be called") },
		db.NewAdminDBForTesting(),
	)

	if len(responder) != 1 {
		t.Fatalf("Want 1 response, got %v", len(responder))
	}
	if responder[0].message != goAway {
		t.Fatalf("Response message: want %v, got %v", goAway, responder[0].message)
	}
}

func TestSkipMessagesToBotWithoutText(t *testing.T) {
	var responder testTextResponder
	handleMessageToBot(
		newResponseContextStub("Admin", ""),
		&responder,
		func(ctx helpers.ResponseContext) { t.Fatal("commands.Execute was not expected to be called") },
		db.NewAdminDBForTesting(),
	)

	if len(responder) != 0 {
		t.Fatalf("Want 0 responses, got %v", len(responder))
	}
}

func TestExecuteCommandInMessageToBot(t *testing.T) {
	executeCommandCalled := false
	var responder testTextResponder
	responseContext := newResponseContextStub("Admin", "/start")
	responseContext.Message.Entities = append(responseContext.Message.Entities, tgbotapi.MessageEntity{Offset: 0, Type: "bot_command"})
	handleMessageToBot(
		responseContext,
		&responder,
		func(ctx helpers.ResponseContext) {
			executeCommandCalled = true
		},
		db.NewAdminDBForTesting(),
	)

	if !executeCommandCalled {
		t.Fatal("executeCommand was not called")
	}
}

func TestBanChannelOfForwardedMessage(t *testing.T) {
	var responder testTextResponder
	responseContext := newResponseContextStub("Admin", "Crappy news!!!")
	responseContext.Message.ForwardFromChat = &tgbotapi.Chat{ID: 666, Title: "Crappy channel"}
	handleMessageToBot(
		responseContext,
		&responder,
		func(ctx helpers.ResponseContext) {
			t.Fatal("commands.Execute was not expected to be called")
		},
		db.NewAdminDBForTesting(),
	)

	if !db.IsChannelIdBanned(666) {
		t.Fatal("Channel was not banned")
	}
	if len(responder) != 1 {
		t.Fatalf("Want 1 response, got %v", len(responder))
	}
	response := responder[0].message
	if response != channelBanResponseText {
		t.Fatalf("Want response %#v, got %#v", channelBanResponseText, response)
	}
}

func TestSkipNonForwardedTextMessagesToBot(t *testing.T) {
	var responder testTextResponder
	handleMessageToBot(
		newResponseContextStub("Admin", "text"),
		&responder,
		func(ctx helpers.ResponseContext) {
			t.Fatal("commands.Execute was not expected to be called")
		},
		db.NewAdminDBForTesting(),
	)
}

func newResponseContextStub(senderUserName, text string) helpers.ResponseContext {
	return helpers.ResponseContext{Message: &tgbotapi.Message{
		From:    &tgbotapi.User{UserName: senderUserName},
		Text:    text,
		Caption: "",
	}}
}
