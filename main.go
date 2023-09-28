package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	readBannedChannelsDatabase()
	readAdminsDatabase()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		message := update.Message
		if message == nil {
			continue
		}

		fmt.Printf("%+v\n", message.Chat)

		if message.IsCommand() && message.Chat.Type == "private" {
			if !isAdmin(message.From.UserName) {
				rejectCommandFromNonAdmin(bot, message)
				continue
			}
			switch message.Command() {
			case "start":
				sendHelp(bot, message)
			case "list":
				listBannedChannels(bot, message)
			case "clear":
				clearBannedChannelsCommand(bot, message)
			default:
				sendUnknownCommandResponse(bot, message)
			}
		} else {
			forwardChat := message.ForwardFromChat
			if forwardChat == nil {
				continue
			}

			log.Printf("Got forward from channel %+v\n", message.ForwardFromChat)

			if message.Chat.Type == "private" {
				channelRecord := channel{Id: message.ForwardFromChat.ID, Title: message.ForwardFromChat.Title}
				banChannel(channelRecord)
				sendBanResponse(bot, message)
			} else {
				if isChannelIdBanned(forwardChat.ID) {
					removeMessage(bot, update.Message)
					mockSender(bot, update.Message)
				}
			}

		}
	}
}

func rejectCommandFromNonAdmin(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	sendWithErrorLogging(
		bot,
		tgbotapi.NewMessage(message.Chat.ID, "Исчезни, я тебя не знаю."),
	)
}

func clearBannedChannelsCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	clearBannedChannels()
	response := tgbotapi.NewMessage(message.Chat.ID, "_Список забаненных каналов очищен._")
	response.ParseMode = "markdown"
	sendWithErrorLogging(bot, response)
}

type channel struct {
	Id    int64
	Title string
}

func listBannedChannels(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	bannedChannels := getBannedChannels()
	text := listChannelsToString(bannedChannels)
	response := tgbotapi.NewMessage(message.Chat.ID, text)
	response.ParseMode = "markdown"
	sendWithErrorLogging(bot, response)
}

func listChannelsToString(bannedChannels []channel) string {
	if len(bannedChannels) == 0 {
		return "_Каналов нет_"
	}
	lines := make([]any, len(bannedChannels))
	for i := 0; i < len(bannedChannels); i++ {
		lines[i] = fmt.Sprintf("%d. %s\n", i+1, bannedChannels[i].Title)
	}
	return fmt.Sprint(lines...)
}

func sendUnknownCommandResponse(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	text := fmt.Sprintf("*Неизвестная команда: `%s`*", message.Command())
	response := tgbotapi.NewMessage(message.Chat.ID, text)
	response.ParseMode = "markdown"
	sendWithErrorLogging(bot, response)
}

func sendBanResponse(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	text := fmt.Sprintf("Канал *%s* забанен.", message.ForwardFromChat.Title)
	response := tgbotapi.NewMessage(message.Chat.ID, text)
	response.ParseMode = "markdown"
	sendWithErrorLogging(bot, response)
}

func sendHelp(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	sendWithErrorLogging(
		bot,
		tgbotapi.NewMessage(
			message.Chat.ID,
			"Этот бот удаляет сообщения, пересланные из забаненных вами каналов.\n\n"+
				"Для того, чтобы забанить канал, перешлите из него сообщение сюда.\n"+
				"Чтобы очистить список забаненных каналов, выполните /clear.\n"+
				"Посмотреть список забаненных каналов — /list.",
		),
	)
}

func removeMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	sendWithErrorLogging(bot, tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID))
}

func sendWithErrorLogging(bot *tgbotapi.BotAPI, c tgbotapi.Chattable) {
	if _, err := bot.Send(c); err != nil {
		log.Println(err.Error())
	}
}

func mockSender(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	text := fmt.Sprintf("%s, вспышка слева!", message.From.FirstName)
	response := tgbotapi.NewMessage(message.Chat.ID, text)
	sendWithErrorLogging(bot, response)
}
