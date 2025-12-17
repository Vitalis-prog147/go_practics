package format

import (
	"fmt"
	"strings"
	"telegram_bot_todo/models"
)

func FormatScoreboard(players []models.PlayerScore) string {
	if len(players) == 0 {
		return ""
	}

	var b strings.Builder
	for i, player := range players {
		position := i + 1
		medal := GetMedalEmoji(position)
		b.WriteString(fmt.Sprintf("%s %d. %s: %d\n", medal, position, player.Name, player.Score))
	}

	return b.String()
}
