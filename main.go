package main

import (
	"log"
	config "telegram_bot_todo/config"
	"telegram_bot_todo/bot"
)

func main() {
	cfg := config.Load()
	
	b, err := bot.New(cfg.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	if err := b.Start(); err != nil {
		log.Panic(err)
	}
}
