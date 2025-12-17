package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram_bot_todo/interfaces"
)

type TelegramMessenger struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramMessenger(bot *tgbotapi.BotAPI) interfaces.Messenger {
	return &TelegramMessenger{bot: bot}
}

func (m *TelegramMessenger) SendMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := m.bot.Send(msg)
	return err
}
