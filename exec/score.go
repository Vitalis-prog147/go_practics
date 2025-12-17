package exec

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram_bot_todo/models"
)

func HandleScore(bot *tgbotapi.BotAPI, chatID int64, teamScore *models.TeamScore) {
	scores := teamScore.GetScores()
	if scores == "" {
		msg := tgbotapi.NewMessage(chatID, "Пока нет очков. Используй /add [имя] чтобы добавить!")
		bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(chatID, "📊 Таблица очков:\n"+scores)
		bot.Send(msg)
	}
}
