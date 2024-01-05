package commands

import (
	"math/rand"
	"strings"
	"time"

	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func init() {
	registerCommand(
		newCommand(
			"stfu",
			"Отправляет пользователя с указанным ником на принудительный отдых, "+
				"запрещая писать ему в беседу на некоторое время",
			func(ctx helpers.ResponseContext) {
				userName := ctx.Message.CommandArguments()
				if len(userName) == 0 || strings.Contains(userName, " ") {
					ctx.SendSilentMarkdownFmt("_Нужен один аргумент — ник пользователя._")
					return
				}

				db.GetMuteDB().MuteUser(userName, randomMuteDuration())
				ctx.SendSilentMarkdownFmt("Пользователь с ником %s отправлен на принудительный отдых. "+
					"Если захочется скостить срок — используй /unstfu.", userName)
			},
		),
	)
}

const minMuteDuration = time.Minute * 10
const maxMuteDuration = time.Minute * 60

func randomMuteDuration() time.Duration {
	return time.Duration(rand.Int63n(int64(maxMuteDuration) - int64(minMuteDuration) + int64(minMuteDuration)))
}
