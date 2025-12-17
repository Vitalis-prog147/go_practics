package models

import (
	"sort"
	"time"
)

type PlayerScore struct {
	Name  string
	Score int
}

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

func (t *TeamScore) GetScores() []PlayerScore {
	t.checkAndResetIfNewDay()

	if len(t.Score) == 0 {
		return nil
	}

	var players []PlayerScore
	for name, score := range t.Score {
		players = append(players, PlayerScore{Name: name, Score: score})
	}

	sort.Slice(players, func(i, j int) bool {
		if players[i].Score == players[j].Score {
			return players[i].Name < players[j].Name
		}
		return players[i].Score > players[j].Score
	})

	return players
}

func (t *TeamScore) HasScores() bool {
	t.checkAndResetIfNewDay()
	return len(t.Score) > 0
}
