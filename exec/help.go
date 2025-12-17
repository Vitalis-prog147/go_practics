package exec

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram_bot_todo/commands"
)

func HandleHelp(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, commands.GetAllCommandsList())
	bot.Send(msg)
}
