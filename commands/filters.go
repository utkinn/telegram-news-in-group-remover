package commands

import (
	"fmt"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/filters"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
	"strings"
)

var filtersCommand = newCommand("filters", "Отображает состояние фильтров.", func(ctx helpers.ResponseContext) {
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
		lines[i] = fmt.Sprintf("%v `[%v]` %v\n      %v", stateEmoji, desc.ID, desc.Name, desc.Desc)
	}
	ctx.SendSilentMarkdownFmt(strings.Join(lines, "\n\n"))
})
