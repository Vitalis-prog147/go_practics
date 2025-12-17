package models

import (
	"fmt"
	"sort"
	"strings"
	"telegram_bot_todo/format"
	
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
		medal := format.GetMedalEmoji(position)
		b.WriteString(fmt.Sprintf("%s %d. %s: %d\n", medal, position, player.name, player.score))
	}

	return b.String()
}

func (t *TeamScore) HasScores() bool {
	t.checkAndResetIfNewDay()
	return len(t.Score) > 0
}
