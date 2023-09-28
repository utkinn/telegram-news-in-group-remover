package main

import (
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

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		handleUpdate(update, bot)
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

	handleMessageToChannel(ctx, update)
}

func handleMessageToBot(ctx helpers.ResponseContext) {
	if !db.IsAdmin(ctx.Message.From.UserName) {
		ctx.SendMarkdownFmt("Исчезни, я тебя не знаю.")
		return
	}

	if ctx.Message.IsCommand() {
		handleCommand(ctx)
	} else {
		banChannelOfForwardedMessage(ctx)
	}
}

func handleCommand(ctx helpers.ResponseContext) {
	switch ctx.Message.Command() {
	case "start":
		commands.Start(ctx)
	case "list":
		commands.List(ctx)
	case "clear":
		commands.Clear(ctx)
	default:
		commands.Unknown(ctx)
	}
}

func handleMessageToChannel(ctx helpers.ResponseContext, update tgbotapi.Update) {
	if db.IsChannelIdBanned(ctx.Message.ForwardFromChat.ID) {
		removeMessage(ctx)
		mockSender(ctx)
	}
}

func banChannelOfForwardedMessage(ctx helpers.ResponseContext) {
	channelRecord := db.Channel{Id: ctx.Message.ForwardFromChat.ID, Title: ctx.Message.ForwardFromChat.Title}
	db.BanChannel(channelRecord)
	sendBanResponse(ctx)
}

func sendBanResponse(ctx helpers.ResponseContext) {
	ctx.SendMarkdownFmt("Канал *%s* забанен.", ctx.Message.ForwardFromChat.Title)
}

func removeMessage(ctx helpers.ResponseContext) {
	helpers.Send(ctx.Bot, tgbotapi.NewDeleteMessage(ctx.Message.Chat.ID, ctx.Message.MessageID))
}

func mockSender(ctx helpers.ResponseContext) {
	ctx.SendMarkdownFmt("%s, вспышка слева!", ctx.Message.From.FirstName)
}
