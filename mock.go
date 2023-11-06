package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	tgbotapi "github.com/utkinn/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
	"math/rand"
	"time"
)

var lastMockAt time.Time

func mockSender(bot *tgbotapi.BotAPI, groupChatId int64, newsSender *tgbotapi.User) {
	if time.Now().Sub(lastMockAt).Minutes() < 1 {
		return
	}
	lastMockAt = time.Now()

	stickerMessage := sendMockSticker(bot, groupChatId)
	mockTextMessage := sendMockTextMessage(bot, groupChatId, newsSender)
	mockCleanupQueue <- mock{messages: []*tgbotapi.Message{stickerMessage, mockTextMessage}, time: time.Now()}
}

func sendMockSticker(bot *tgbotapi.BotAPI, groupChatId int64) *tgbotapi.Message {
	stickerMessageRequest := tgbotapi.NewSticker(groupChatId, db.GetRandomMockStickerFileId())
	stickerMessageRequest.DisableNotification = true
	return helpers.Send(bot, stickerMessageRequest)
}

func sendMockTextMessage(bot *tgbotapi.BotAPI, groupChatId int64, newsSender *tgbotapi.User) *tgbotapi.Message {
	senderFunnyName := db.GetNameForUser(newsSender)
	mockTextMessageRequest := tgbotapi.NewMessage(groupChatId, fmt.Sprintf("%s, вспышка слева!", senderFunnyName))
	userNameHash := sha256.Sum256([]byte(newsSender.UserName))
	if hex.EncodeToString(userNameHash[:]) == "2d2aa474c3574e0c36d120d1a60f8f729fc355b8ac379c3cb529609ee60788f2" {
		mockTextMessageRequest.Text = fmt.Sprintf(
			"%s, [иди поищи работу](https://magnitogorsk.hh.ru/search/vacancy?L_save_area=true&text=&excluded_text=&area=1399&salary=&currency_code=RUR&experience=noExperience&employment=full&employment=part&schedule=fullDay&schedule=shift&schedule=flexible&order_by=relevance&search_period=0&items_on_page=50)",
			senderFunnyName,
		)
		mockTextMessageRequest.ParseMode = tgbotapi.ModeMarkdown
	}
	if rand.Intn(100) <= 2 {
		mockTextMessageRequest.Text = fmt.Sprintf("%v, вспышка _сверху_ 💥🔝 !", senderFunnyName)
		mockTextMessageRequest.ParseMode = tgbotapi.ModeMarkdown
	}
	return helpers.Send(bot, mockTextMessageRequest)
}
