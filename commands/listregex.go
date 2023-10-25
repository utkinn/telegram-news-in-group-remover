package commands

import (
	"fmt"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
	"strings"
)

var listRegexCommand = newCommand("listregex", "Пристальный присмотр - показать список запрещенных регулярных выражений", func(ctx helpers.ResponseContext) {
	regexes := db.GetRegexes()
	if len(regexes) == 0 {
		ctx.SendSilentMarkdownFmt("_Нет регулярных выражений_")
		return
	}
	regexesOrderedListLines := make([]string, len(regexes))
	for i, regex := range regexes {
		regexesOrderedListLines[i] = fmt.Sprintf("%d. %s", i+1, regex)
	}
	ctx.SendSilentFmt(strings.Join(regexesOrderedListLines, "\n"))
})
