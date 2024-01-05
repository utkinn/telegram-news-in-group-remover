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
	filterID := ctx.Message.CommandArguments()
	if len(filterID) == 0 || strings.Contains(filterID, " ") {
		ctx.SendSilentMarkdownFmt("_Нужен один аргумент — ID фильтра." +
			" Он указывается в квадратных скобках в выводе /filters._")
		return
	}

	if !filters.ValidID(filterID) {
		ctx.SendSilentMarkdownFmt("_Неверный ID фильтра. Сверьтесь с выводом команды /filters._")
		return
	}

	db.GetFilterToggleDB().SetFilterEnabled(filterID, newState)
	ctx.SendSilentMarkdownFmt("_Состояние фильтра обновлено._")
}
