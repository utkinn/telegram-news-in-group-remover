package msgremoval

import (
	"fmt"
	"math/rand"
	"time"

	tgbotapi "github.com/utkinn/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

var lastMockTimestampsPerChatId = make(map[int64]time.Time)

func MockUser(bot *tgbotapi.BotAPI, groupChatId int64, user *tgbotapi.User) {
	lastMockAt, ok := lastMockTimestampsPerChatId[groupChatId]
	if ok && time.Since(lastMockAt).Minutes() < 1 {
		return
	}
	lastMockTimestampsPerChatId[groupChatId] = time.Now()

	stickerMessage := sendMockSticker(bot, groupChatId)
	mockTextMessage := sendMockTextMessage(bot, groupChatId, user)
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
	mockTextMessageRequest.DisableNotification = true
	if rand.Intn(100) <= 2 {
		mockTextMessageRequest.Text = fmt.Sprintf("%v, вспышка _сверху_ 💥🔝 !", senderFunnyName)
		mockTextMessageRequest.ParseMode = tgbotapi.ModeMarkdown
	}
	return helpers.Send(bot, mockTextMessageRequest)
}
