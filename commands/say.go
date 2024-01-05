package commands

import (
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func init() {
	registerCommand(
		newCommand(
			"say",
			"<номер чата> <текст> — Брякнуть что-нибудь в чат работягам",
			func(ctx helpers.ResponseContext) {
				chatNumStrAndText := strings.SplitN(ctx.Message.CommandArguments(), " ", 2)
				if len(chatNumStrAndText) != 2 {
					ctx.SendSilentFmt("Ты не указал текст или номер чата. " +
						"Правильный формат:\n/say <номер чата> <текст>\nНомер чата можешь узнать командой /chats.")
					return
				}

				chatNumStr := chatNumStrAndText[0]
				text := chatNumStrAndText[1]

				chatNum, _ := strconv.Atoi(chatNumStr)
				chatID, ok := db.GetChatsDB().GetIDByOrdinal(chatNum)
				if !ok {
					ctx.SendSilentFmt("Неправильный номер чата: %s\n\n"+
						"Используй /chats, чтобы посмотреть, куда можно писать", chatNumStr)
					return
				}

				msg := tgbotapi.NewMessage(chatID, text)
				msg.DisableNotification = true
				copyMarkupFromTextCmdArg(*ctx.Message, &msg, len(text))
				helpers.Send(ctx.Bot, msg)
			}),
	)
}
