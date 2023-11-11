package commands

import (
	_ "embed"

	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

//go:embed start.md
var startText string

func init() {
	registerCommand(
		newCommand("start", "Справка", func(ctx helpers.ResponseContext) {
			ctx.SendSilentMarkdownFmt(startText)
		}),
	)
}
