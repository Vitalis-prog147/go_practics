package exec

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
	"telegram_bot_todo/commands"
)

func HandleUnknown(bot *tgbotapi.BotAPI, chatID int64, text string) {
	if strings.HasPrefix(text, "/") {
		msg := tgbotapi.NewMessage(chatID,
			fmt.Sprintf("❌ Неизвестная команда: %s\n\n%s",
				text, commands.GetAllCommandsList()))
		bot.Send(msg)
	}
}
