package msgremoval

import (
	"fmt"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

var lastMockTimestampsPerChatID = make(map[int64]time.Time)

func MockUser(bot *tgbotapi.BotAPI, groupChatID int64, user *tgbotapi.User) {
	lastMockAt, ok := lastMockTimestampsPerChatID[groupChatID]
	if ok && time.Since(lastMockAt).Minutes() < 1 {
		return
	}

	lastMockTimestampsPerChatID[groupChatID] = time.Now()

	stickerMessage := sendMockSticker(bot, groupChatID)
	mockTextMessage := sendMockTextMessage(bot, groupChatID, user)
	mockCleanupQueue <- mock{messages: []*tgbotapi.Message{stickerMessage, mockTextMessage}, time: time.Now()}
}

func sendMockSticker(bot *tgbotapi.BotAPI, groupChatID int64) *tgbotapi.Message {
	stickerMessageRequest := tgbotapi.NewSticker(groupChatID, db.GetStickerDB().GetRandomMockStickerFileID())
	stickerMessageRequest.DisableNotification = true

	return helpers.Send(bot, stickerMessageRequest)
}

func sendMockTextMessage(bot *tgbotapi.BotAPI, groupChatID int64, newsSender *tgbotapi.User) *tgbotapi.Message {
	senderFunnyName := db.GetNameReplacementDB().GetNameForUser(newsSender)
	mockTextMessageRequest := tgbotapi.NewMessage(groupChatID, fmt.Sprintf("%s, вспышка слева!", senderFunnyName))
	mockTextMessageRequest.DisableNotification = true

	if rand.Intn(100) <= 2 { //nolint:gomnd
		mockTextMessageRequest.Text = fmt.Sprintf("%v, вспышка _сверху_ 💥🔝 !", senderFunnyName)
		mockTextMessageRequest.ParseMode = tgbotapi.ModeMarkdown
	}

	return helpers.Send(bot, mockTextMessageRequest)
}
