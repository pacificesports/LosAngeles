package model

import (
	"encoding/json"
	"time"
)

type TournamentTeam struct {
	TournamentID string          `gorm:"primaryKey" json:"tournament_id"`
	TeamID       string          `gorm:"primaryKey" json:"team_id"`
	Team         json.RawMessage `gorm:"-" json:"team"`
	CreatedAt    time.Time       `gorm:"autoCreateTime" json:"created_at"`
}

func (TournamentTeam) TableName() string {
	return "tournament_team"
}
