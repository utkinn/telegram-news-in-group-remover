package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/utkinn/telegram-news-in-group-remover/db"
	"github.com/utkinn/telegram-news-in-group-remover/helpers"
)

type commandCallback func(ctx helpers.ResponseContext)

type command struct {
	name, help string
	callback   commandCallback
	hidden     bool
}

func newCommand(name, help string, callback commandCallback) command {
	return command{
		name:     name,
		help:     help,
		callback: callback,
	}
}

func newSuperAdminCommand(name, help string, callback commandCallback) command {
	return command{
		name: name,
		help: help,
		callback: func(ctx helpers.ResponseContext) {
			if !db.GetAdminDB().IsSuperAdmin(ctx.Message.From.UserName) {
				ctx.SendSilentMarkdownFmt("_Эта команда доступна только для суперадмина._")
				return
			}
			callback(ctx)
		},
	}
}

func newHiddenSuperAdminCommand(name string, callback commandCallback) command {
	return command{
		name:   name,
		hidden: true,
		callback: func(ctx helpers.ResponseContext) {
			if !db.GetAdminDB().IsSuperAdmin(ctx.Message.From.UserName) {
				ctx.SendSilentMarkdownFmt("_Эта команда доступна только для суперадмина._")
				return
			}
			callback(ctx)
		},
	}
}

var commands []command

func registerCommand(cmd command) {
	commands = append(commands, cmd)
}

func GetCommandList() []tgbotapi.BotCommand {
	cmdList := make([]tgbotapi.BotCommand, 0, len(commands))
	for _, cmd := range commands {
		if !cmd.hidden {
			cmdList = append(cmdList, tgbotapi.BotCommand{Command: cmd.name, Description: cmd.help})
		}
	}
	return cmdList
}

func Execute(ctx helpers.ResponseContext) {
	cmdName := ctx.Message.Command()
	for _, cmd := range commands {
		if cmd.name == cmdName {
			cmd.callback(ctx)
			return
		}
	}
	ctx.SendSilentMarkdownFmt("_Неизвестная команда: %s_", cmdName)
}

func copyMarkupFromTextCmdArg(from tgbotapi.Message, to *tgbotapi.MessageConfig, textLength int) {
	to.Entities = make([]tgbotapi.MessageEntity, 0, len(from.Entities)-1) // -1 for bot_command entity
	for _, ent := range from.Entities {
		if ent.Type != "bot_command" {
			ent.Offset -= len(from.Text) - textLength
			to.Entities = append(to.Entities, ent)
		}
	}
}
