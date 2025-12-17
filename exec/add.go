package exec

import (
	"fmt"
	"strings"
	"telegram_bot_todo/interfaces"
	"telegram_bot_todo/models"
)

func HandleAdd(messenger interfaces.Messenger, chatID int64, text string, teamScore *models.TeamScore) error {
	parts := strings.Fields(text)
	if len(parts) < 2 {
		return messenger.SendMessage(chatID, "❌ Использование: /add [имя]\nПример: /add Иван")
	}

	name := strings.Join(parts[1:], " ")
	teamScore.AddScore(name)
	return messenger.SendMessage(chatID, fmt.Sprintf("✅ Очко добавлено игроку: %s", name))
}
