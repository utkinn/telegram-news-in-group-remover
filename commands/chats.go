package commands

import (
	"fmt"

	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func init() {
	registerCommand(
		newCommand("chats", "Список групповых чатов, куда можно писать командой /say", func(ctx helpers.ResponseContext) {
			chats := db.GetChatsDB().Get()
			if len(chats) == 0 {
				ctx.SendSilentFmt("Нет ни одного чата, куда можно писать командой /say. Подожди, пока в один из чатов, где есть я, кто-нибудь напишет.")
				return
			}

			var text string
			for i, chat := range chats {
				text += fmt.Sprintf("%d. %s\n", i+1, chat.Title)
			}
			ctx.SendSilentFmt(text)
		}),
	)
}
