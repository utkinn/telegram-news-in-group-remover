package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/utkinn/telegram-news-in-group-remover/commands"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func main() {
	db.Load()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	setUpCommandList(bot)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		handleUpdate(update, bot)
	}
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
		handleMessageToBot(ctx)
		return
	}

	handleMessageToGroup(ctx)
}

func handleMessageToBot(ctx helpers.ResponseContext) {
	if ctx.Message.Sticker != nil {
		fmt.Println(ctx.Message.Sticker.FileID)
	}

	if !db.IsAdmin(ctx.Message.From.UserName) {
		ctx.SendSilentMarkdownFmt("Исчезни, я тебя не знаю.")
		return
	}

	if ctx.Message.IsCommand() {
		commands.Execute(ctx)
	} else {
		banChannelOfForwardedMessage(ctx)
	}
}

func handleMessageToGroup(ctx helpers.ResponseContext) {
	if !passesScrutinyFilters(ctx.Message) {
		removeMessage(ctx)
		mockSender(ctx.Bot, ctx.Message.Chat.ID, ctx.Message.From)
	}

	if ctx.Message.ForwardFromChat == nil {
		return
	}
	if db.IsChannelIdBanned(ctx.Message.ForwardFromChat.ID) {
		removeMessage(ctx)
		mockSender(ctx.Bot, ctx.Message.Chat.ID, ctx.Message.From)
	} else {
		forwardMemory = append(forwardMemory, newForwardMemoryItem(ctx.Message))
	}
}

func banChannelOfForwardedMessage(ctx helpers.ResponseContext) {
	if ctx.Message.ForwardFromChat == nil {
		return
	}
	channelRecord := db.Channel{Id: ctx.Message.ForwardFromChat.ID, Title: ctx.Message.ForwardFromChat.Title}
	db.BanChannel(channelRecord)
	removeMessagesFromNewlyBannedChannel(ctx.Bot, channelRecord)
	sendBanResponse(ctx)
}

func removeMessagesFromNewlyBannedChannel(bot *tgbotapi.BotAPI, chanRec db.Channel) {
	userNamesMockedSoFar := map[string]bool{}
	for _, item := range forwardMemory {
		if item.channelId != chanRec.Id {
			continue
		}

		helpers.Send(bot, tgbotapi.NewDeleteMessage(item.groupChatId, item.messageId))
		if !userNamesMockedSoFar[item.from.UserName] {
			mockSender(bot, item.groupChatId, &item.from)
			userNamesMockedSoFar[item.from.UserName] = true
		}
	}
}

func sendBanResponse(ctx helpers.ResponseContext) {
	ctx.SendSilentMarkdownFmt("Канал *%s* забанен.", ctx.Message.ForwardFromChat.Title)
}

func removeMessage(ctx helpers.ResponseContext) {
	helpers.Send(ctx.Bot, tgbotapi.NewDeleteMessage(ctx.Message.Chat.ID, ctx.Message.MessageID))
}

func mockSender(bot *tgbotapi.BotAPI, groupChatId int64, newsSender *tgbotapi.User) {
	message := tgbotapi.NewSticker(groupChatId, db.GetRandomMockStickerFileId())
	message.DisableNotification = true
	helpers.Send(bot, message)
	msg := tgbotapi.NewMessage(groupChatId, fmt.Sprintf("%s, вспышка слева!", db.GetNameForUser(newsSender)))
	msg.DisableNotification = true
	helpers.Send(bot, msg)
}
