package exec

import (
	"fmt"
	"strings"
	"telegram_bot_todo/commands"
	"telegram_bot_todo/interfaces"
)

func HandleUnknown(messenger interfaces.Messenger, chatID int64, text string) error {
	if strings.HasPrefix(text, "/") {
		return messenger.SendMessage(chatID,
			fmt.Sprintf("❌ Неизвестная команда: %s\n\n%s",
				text, commands.GetAllCommandsList()))
	}
	return nil
}
