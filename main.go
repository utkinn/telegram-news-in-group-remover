package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

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

	notifyRestart(bot)

	setUpCommandList(bot)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		handleUpdate(update, bot)
	}
}

func notifyRestart(bot *tgbotapi.BotAPI) {
	superAdminChatId := db.GetSuperAdminChatId()
	if superAdminChatId == db.SuperAdminChatIdNotSet {
		return
	}

	msg := tgbotapi.NewMessage(superAdminChatId, "_Бот пезерапущен_")
	msg.ParseMode = "markdown"
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

func handleMessageToGroup(ctx helpers.ResponseContext) {
	if !passesScrutinyFilters(ctx.Message) {
		removeMessageAndMockSender(ctx.Bot, ctx.Message)
	}

	if ctx.Message.ForwardFromChat == nil {
		return
	}
	if db.IsChannelIdBanned(ctx.Message.ForwardFromChat.ID) {
		removeMessageAndMockSender(ctx.Bot, ctx.Message)
	} else {
		forwardMemory = append(forwardMemory, newForwardMemoryItem(ctx.Message))
	}
}

func removeMessageAndMockSender(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	removeMessage(bot, message)
	mockSender(bot, message.Chat.ID, message.From)
}

func banChannelOfForwardedMessage(ctx helpers.ResponseContext) {
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

var lastMockAt time.Time

func mockSender(bot *tgbotapi.BotAPI, groupChatId int64, newsSender *tgbotapi.User) {
	if time.Now().Sub(lastMockAt).Minutes() < 1 {
		return
	}
	lastMockAt = time.Now()
	message := tgbotapi.NewSticker(groupChatId, db.GetRandomMockStickerFileId())
	message.DisableNotification = true
	helpers.Send(bot, message)
	msg := tgbotapi.NewMessage(groupChatId, fmt.Sprintf("%s, вспышка слева!", db.GetNameForUser(newsSender)))
	userNameHash := sha256.Sum256([]byte(newsSender.UserName))
	if hex.EncodeToString(userNameHash[:]) == "2d2aa474c3574e0c36d120d1a60f8f729fc355b8ac379c3cb529609ee60788f2" {
		msg.Text = fmt.Sprintf(
			"%s, [иди поищи работу](https://magnitogorsk.hh.ru/search/vacancy?L_save_area=true&text=&excluded_text=&area=1399&salary=&currency_code=RUR&experience=noExperience&employment=full&employment=part&schedule=fullDay&schedule=shift&schedule=flexible&order_by=relevance&search_period=0&items_on_page=50)",
			db.GetNameForUser(newsSender),
		)
		msg.ParseMode = "markdown"
	}
	msg.DisableNotification = true
	helpers.Send(bot, msg)
}
