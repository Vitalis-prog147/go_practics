package format

func GetMedalEmoji(position int) string {
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
