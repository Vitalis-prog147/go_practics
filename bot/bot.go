package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram_bot_todo/commands"
	"telegram_bot_todo/exec"
	"telegram_bot_todo/infrastructure/telegram"
	"telegram_bot_todo/interfaces"
	"telegram_bot_todo/models"
)

type Bot struct {
	api       *tgbotapi.BotAPI
	messenger interfaces.Messenger
	teamScore *models.TeamScore
}

func New(token string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании бота: %w", err)
	}

	return &Bot{
		api:       api,
		messenger: telegram.NewTelegramMessenger(api),
		teamScore: models.NewTeamScore(),
	}, nil
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.api.GetUpdatesChan(u)

	fmt.Println("🤖 Бот запущен и ждет команд!")

	for update := range updates {
		if update.Message == nil {
			continue
		}

		b.handleUpdate(update)
	}

	return nil
}

func (b *Bot) handleUpdate(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	text := update.Message.Text

	command := commands.ParseCommand(text)

	switch command {
	case commands.CommandStart:
		exec.HandleStart(b.messenger, chatID)

	case commands.CommandScore:
		exec.HandleScore(b.messenger, chatID, b.teamScore)

	case commands.CommandAdd:
		exec.HandleAdd(b.messenger, chatID, text, b.teamScore)

	case commands.CommandHelp:
		exec.HandleHelp(b.messenger, chatID)

	default:
		exec.HandleUnknown(b.messenger, chatID, text)
	}
}

func (b *Bot) GetAPI() *tgbotapi.BotAPI {
	return b.api
}
