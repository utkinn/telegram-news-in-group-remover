package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	tgbotapi "github.com/utkinn/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/commands"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/filters"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
	"log"
	"os"
)

var mockCleanupQueue = make(chan mock, 100)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	filters.InitFilters(bot)

	notifyRestart(bot)

	setUpCommandList(bot)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	go mockCleaner(bot)

	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		handleUpdate(update, bot)
	}
}

func notifyRestart(bot *tgbotapi.BotAPI) {
	if _, skipSet := os.LookupEnv("SKIP_RESTART_NOTIFICATION"); skipSet {
		return
	}

	superAdminChatId := db.GetSuperAdminChatId()
	if superAdminChatId == db.SuperAdminChatIdNotSet {
		return
	}

	msg := tgbotapi.NewMessage(superAdminChatId, "_Бот пезерапущен_")
	msg.ParseMode = tgbotapi.ModeMarkdown
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

	if ctx.Message.Caption == "" && ctx.Message.Text == "" {
		return
	}

	if ctx.Message.IsCommand() {
		commands.Execute(ctx)
	} else {
		banChannelOfForwardedMessage(ctx)
	}
}

var offendingMediaGroupId string

func handleMessageToGroup(ctx helpers.ResponseContext) {
	if ctx.Message.MediaGroupID != "" && ctx.Message.MediaGroupID == offendingMediaGroupId {
		removeMessage(ctx.Bot, ctx.Message)
	}

	// Get ready to remove the entire album
	offendingMediaGroupId = ctx.Message.MediaGroupID

	if allowed, shouldSuppressMock := filters.IsMessageAllowed(ctx.Message); !allowed {
		removeMessage(ctx.Bot, ctx.Message)
		if !shouldSuppressMock {
			mockSender(ctx.Bot, ctx.Message.Chat.ID, ctx.Message.From)
		}
	} else {
		if ctx.Message.ForwardFromChat != nil {
			forwardMemory = append(forwardMemory, newForwardMemoryItem(ctx.Message))
		}
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

func removeMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	helpers.Send(bot, tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID))
	db.RecordMessageRemoval(message)
}
