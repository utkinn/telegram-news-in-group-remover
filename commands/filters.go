package commands

import (
	"fmt"
	"strings"

	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/filters"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func init() {
	registerCommand(
		newCommand("filters", "Отображает состояние фильтров.", func(ctx helpers.ResponseContext) {
			filterList := filters.List()
			lines := make([]string, len(filterList))
			for i, filter := range filterList {
				desc := filter.Description()

				var stateEmoji string
				if db.IsFilterEnabled(desc.ID) {
					stateEmoji = "🟢"
				} else {
					stateEmoji = "🔴"
				}

				var scrutinyNotice string
				if filter.ScrutinyModeOnly() {
					scrutinyNotice = "\n      ▪️_Только для пристального присмотра (/scrutiny)_"
				}

				lines[i] = fmt.Sprintf("%v `[%v]` %v\n      %v%v", stateEmoji, desc.ID, desc.Name, desc.Desc, scrutinyNotice)
			}
			ctx.SendSilentMarkdownFmt(strings.Join(lines, "\n\n"))
		}),
	)
}
