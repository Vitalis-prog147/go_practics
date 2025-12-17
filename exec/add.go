package exec

import (
	"fmt"
	"strings"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram_bot_todo/models"
)

func HandleAdd(bot *tgbotapi.BotAPI, chatID int64, text string, teamScore *models.TeamScore) bool {
	parts := strings.Fields(text)
	if len(parts) < 2 {
		msg := tgbotapi.NewMessage(chatID, "❌ Использование: /add [имя]\nПример: /add Иван")
		bot.Send(msg)
		return false
	}

	name := strings.Join(parts[1:], " ")
	teamScore.AddScore(name)
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("✅ Очко добавлено игроку: %s", name))
	bot.Send(msg)
	return true
}
