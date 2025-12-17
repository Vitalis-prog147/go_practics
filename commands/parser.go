package commands

import (
	"fmt"
	"strings"
)

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
