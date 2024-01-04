package commands

import (
	"strings"

	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/filters"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

func init() {
	registerCommand(
		newCommand("filteron", "<id> — Включает фильтр с указанным ID.", func(ctx helpers.ResponseContext) {
			filterOnOffCallback(ctx, true)
		}),
	)

	registerCommand(
		newCommand("filteroff", "<id> — Выключает фильтр с указанным ID.", func(ctx helpers.ResponseContext) {
			filterOnOffCallback(ctx, false)
		}),
	)
}

func filterOnOffCallback(ctx helpers.ResponseContext, newState bool) {
	id := ctx.Message.CommandArguments()
	if len(id) == 0 || strings.Contains(id, " ") {
		ctx.SendSilentMarkdownFmt("_Нужен один аргумент — ID фильтра. Он указывается в квадратных скобках в выводе /filters._")
		return
	}

	if !filters.ValidID(id) {
		ctx.SendSilentMarkdownFmt("_Неверный ID фильтра. Сверьтесь с выводом команды /filters._")
		return
	}

	db.GetFilterToggleDB().SetFilterEnabled(id, newState)
	ctx.SendSilentMarkdownFmt("_Состояние фильтра обновлено._")
}
