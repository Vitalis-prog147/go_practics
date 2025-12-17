package exec

import (
	"telegram_bot_todo/format"
	"telegram_bot_todo/interfaces"
	"telegram_bot_todo/models"
)

func HandleScore(messenger interfaces.Messenger, chatID int64, teamScore *models.TeamScore) error {
	players := teamScore.GetScores()
	if len(players) == 0 {
		return messenger.SendMessage(chatID, "Пока нет очков. Используй /add [имя] чтобы добавить!")
	}

	formatted := format.FormatScoreboard(players)
	return messenger.SendMessage(chatID, "📊 Таблица очков:\n"+formatted)
}
