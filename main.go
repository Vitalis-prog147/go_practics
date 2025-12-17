package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/joho/godotenv"
	"log"
	"sort"
	"strings"
	config "telegram_bot_todo/config"
	"time"
)

type TeamScore struct {
	Score         map[string]int
	LastResetTime time.Time
}

func NewTeamScore() *TeamScore {
	return &TeamScore{
		Score:         make(map[string]int),
		LastResetTime: time.Now(),
	}
}

func (t *TeamScore) checkAndResetIfNewDay() {
	now := time.Now()
	if now.Year() != t.LastResetTime.Year() ||
		now.Month() != t.LastResetTime.Month() ||
		now.Day() != t.LastResetTime.Day() {
		t.Score = make(map[string]int)
		t.LastResetTime = now
	}
}

func (t *TeamScore) AddScore(name string) {
	t.checkAndResetIfNewDay()
	t.Score[name]++
}

func getMedalEmoji(position int) string {
	switch position {
	case 1:
		return "🥇"
	case 2:
		return "🥈"
	case 3:
		return "🥉"
	default:
		emojis := []string{"🏅", "⭐", "💫", "✨", "🌟", "🎖️", "🎯", "🔥"}
		if position-1 < len(emojis) {
			return emojis[position-1]
		}
		return "📍"
	}
}

func (t *TeamScore) GetScores() string {
	t.checkAndResetIfNewDay()

	if len(t.Score) == 0 {
		return ""
	}

	type playerScore struct {
		name  string
		score int
	}

	var players []playerScore
	for name, score := range t.Score {
		players = append(players, playerScore{name: name, score: score})
	}

	sort.Slice(players, func(i, j int) bool {
		if players[i].score == players[j].score {
			return players[i].name < players[j].name
		}
		return players[i].score > players[j].score
	})

	var b strings.Builder
	for i, player := range players {
		position := i + 1
		medal := getMedalEmoji(position)
		b.WriteString(fmt.Sprintf("%s %d. %s: %d\n", medal, position, player.name, player.score))
	}

	return b.String()
}

func main() {
	bot, err := tgbotapi.NewBotAPI(config.Load().TelegramBotToken)
	if err != nil {
		log.Panic("Ошибка при создании бота:", err)
	}

	teamScore := NewTeamScore()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)

	fmt.Println("🤖 Бот запущен и ждет команд!")

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		text := update.Message.Text

		command := ParseCommand(text)

		switch command {
		case CommandStart:
			msg := tgbotapi.NewMessage(chatID, "Привет! Я бот для подсчета очков.\n\n"+GetAllCommandsList())
			bot.Send(msg)

		case CommandScore:
			scores := teamScore.GetScores()
			if scores == "" {
				msg := tgbotapi.NewMessage(chatID, "Пока нет очков. Используй /add [имя] чтобы добавить!")
				bot.Send(msg)
			} else {
				msg := tgbotapi.NewMessage(chatID, "📊 Таблица очков:\n"+scores)
				bot.Send(msg)
			}

		case CommandAdd:
			parts := strings.Fields(text)
			if len(parts) < 2 {
				msg := tgbotapi.NewMessage(chatID, "❌ Использование: /add [имя]\nПример: /add Иван")
				bot.Send(msg)
				continue
			}

			name := strings.Join(parts[1:], " ")
			teamScore.AddScore(name)
			msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("✅ Очко добавлено игроку: %s", name))
			bot.Send(msg)

		case CommandHelp:
			msg := tgbotapi.NewMessage(chatID, GetAllCommandsList())
			bot.Send(msg)

		default:
			if strings.HasPrefix(text, "/") {
				msg := tgbotapi.NewMessage(chatID,
					fmt.Sprintf("❌ Неизвестная команда: %s\n\n%s",
						text, GetAllCommandsList()))
				bot.Send(msg)
			}
		}
	}
}
