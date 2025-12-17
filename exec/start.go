package exec

import (
	"telegram_bot_todo/commands"
	"telegram_bot_todo/interfaces"
)

func HandleStart(messenger interfaces.Messenger, chatID int64) error {
	return messenger.SendMessage(chatID, "Привет! Я бот для подсчета очков.\n\n"+commands.GetAllCommandsList())
}
