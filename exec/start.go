package exec

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram_bot_todo/commands"
)

func HandleStart(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Привет! Я бот для подсчета очков.\n\n"+commands.GetAllCommandsList())
	bot.Send(msg)
}
