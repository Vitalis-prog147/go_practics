package main

import (
	"fmt"
	"strings"
)

type CommandType string

const (
	CommandStart CommandType = "/start"
	CommandScore CommandType = "/score"
	CommandAdd   CommandType = "/add"
	CommandHelp  CommandType = "/help"
	CommandNone  CommandType = ""
)

func GetAllCommands() []CommandType {
	return []CommandType{
		CommandStart,
		CommandScore,
		CommandAdd,
		CommandHelp,
	}
}

func (c CommandType) GetDescription() string {
	switch c {
	case CommandStart:
		return "начать работу с ботом"
	case CommandScore:
		return "показать таблицу очков"
	case CommandAdd:
		return "добавить очко игроку (использование: /add [имя])"
	case CommandHelp:
		return "показать справку по командам"
	default:
		return "неизвестная команда - обратитесь в техническую поддержку"
	}
}

func ParseCommand(text string) CommandType {
	text = strings.TrimSpace(strings.ToLower(text))

	if !strings.HasPrefix(text, "/") {
		return CommandNone
	}

	parts := strings.Fields(text)
	if len(parts) == 0 {
		return CommandNone
	}

	command := CommandType(parts[0])

	switch command {
	case CommandStart, CommandScore, CommandAdd, CommandHelp:
		return command
	default:
		return CommandNone
	}
}

func GetAllCommandsList() string {
	var builder strings.Builder
	builder.WriteString("📋 Доступные команды:\n\n")

	commands := GetAllCommands()
	for _, cmd := range commands {
		builder.WriteString(fmt.Sprintf("%s - %s\n", cmd, cmd.GetDescription()))
	}

	return builder.String()
}
