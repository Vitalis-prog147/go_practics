package exec

import (
	"telegram_bot_todo/commands"
	"telegram_bot_todo/interfaces"
)

func HandleHelp(messenger interfaces.Messenger, chatID int64) error {
	return messenger.SendMessage(chatID, commands.GetAllCommandsList())
}
