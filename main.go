package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
	tgbotapi "github.com/utkinn/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/commands"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/filters"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
	"github.com/utkinn/telegram-news-in-group-remover/msgmem"
	"github.com/utkinn/telegram-news-in-group-remover/msgremoval"
)

func main() {
	bot := createBotWithNetworkRetry()

	notifyRestart(bot)

	setUpCommandList(bot)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	go msgremoval.MockCleaner(bot)

	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		handleUpdate(update, bot)
	}
}

func createBotWithNetworkRetry() *tgbotapi.BotAPI {
	var bot *tgbotapi.BotAPI
	err := errors.New("")
	for err != nil {
		bot, err = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
		if err != nil && !strings.Contains(err.Error(), "network is unreachable") {
			panic(err)
		}
		if err != nil {
			time.Sleep(time.Second * 10)
		}
	}
	return bot
}

func notifyRestart(bot *tgbotapi.BotAPI) {
	if _, skipSet := os.LookupEnv("SKIP_RESTART_NOTIFICATION"); skipSet {
		return
	}

	superAdminChatId := db.GetSuperAdminChatId()
	if superAdminChatId == db.SuperAdminChatIdNotSet {
		return
	}

	msg := tgbotapi.NewMessage(superAdminChatId, "_Бот перезапущен_")
	msg.ParseMode = tgbotapi.ModeMarkdown
	msg.DisableNotification = true
	helpers.Send(bot, msg)
}

func setUpCommandList(bot *tgbotapi.BotAPI) {
	commandListConfig := tgbotapi.NewSetMyCommandsWithScope(
		tgbotapi.NewBotCommandScopeAllPrivateChats(),
		commands.GetCommandList()...,
	)
	if _, err := bot.Request(commandListConfig); err != nil {
		log.Printf("Failed to hide commands in groups: %v", err.Error())
	}
}

func handleUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	message := update.Message
	if message == nil {
		return
	}

	ctx := helpers.ResponseContext{Message: message, Bot: bot}

	if message.Chat.Type == "private" {
		handleMessageToBot(ctx, newTelegramTextResponder(bot, message.Chat.ID), commands.Execute)
		return
	}

	handleMessageToGroup(ctx)
}

const goAway = "Исчезни, я тебя не знаю."

func handleMessageToBot(ctx helpers.ResponseContext, resp textResponder, executeCommand func(ctx helpers.ResponseContext)) {
	if ctx.Message.Sticker != nil {
		fmt.Println(ctx.Message.Sticker.FileID)
	}

	if !db.IsAdmin(ctx.Message.From.UserName) {
		resp.RespondTextf(tgbotapi.ModeMarkdown, true, goAway)
		return
	}

	if ctx.Message.Caption == "" && ctx.Message.Text == "" {
		return
	}

	if ctx.Message.IsCommand() {
		executeCommand(ctx)
	} else {
		banChannelOfForwardedMessage(ctx, resp)
	}
}

var offendingMediaGroupId string

func handleMessageToGroup(ctx helpers.ResponseContext) {
	db.AddChat(db.Chat{Id: ctx.Message.Chat.ID, Title: ctx.Message.Chat.Title})

	if ctx.Message.MediaGroupID != "" && ctx.Message.MediaGroupID == offendingMediaGroupId {
		msgremoval.Remove(ctx.Bot, ctx.Message)
	}

	if allowed, shouldSuppressMock := filters.IsMessageAllowed(ctx); !allowed {
		msgremoval.Remove(ctx.Bot, ctx.Message)
		// Get ready to remove the entire album
		offendingMediaGroupId = ctx.Message.MediaGroupID
		if !shouldSuppressMock {
			msgremoval.MockUser(ctx.Bot, ctx.Message.Chat.ID, ctx.Message.From)
		}
	} else {
		msgmem.Add(ctx.Message)
	}
}

func banChannelOfForwardedMessage(ctx helpers.ResponseContext, resp textResponder) {
	if ctx.Message.ForwardFromChat == nil {
		return
	}
	channelRecord := db.Channel{Id: ctx.Message.ForwardFromChat.ID, Title: ctx.Message.ForwardFromChat.Title}
	db.BanChannel(channelRecord)
	removeMessagesFromNewlyBannedChannel(ctx.Bot, channelRecord)
	sendBanResponse(resp)
}

func removeMessagesFromNewlyBannedChannel(bot *tgbotapi.BotAPI, chanRec db.Channel) {
	userNamesMockedSoFar := map[string]bool{}
	for _, item := range msgmem.Get() {
		if item.ForwardFromChat.ID != chanRec.Id {
			continue
		}

		helpers.Send(bot, tgbotapi.NewDeleteMessage(item.Chat.ID, item.MessageID))
		if !userNamesMockedSoFar[item.From.UserName] {
			msgremoval.MockUser(bot, item.Chat.ID, item.From)
			userNamesMockedSoFar[item.From.UserName] = true
		}
	}
}

const channelBanResponseText = "Этот канал забанен."

func sendBanResponse(resp textResponder) {
	resp.RespondTextf("", true, channelBanResponseText)
}
