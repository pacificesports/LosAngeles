package model

import "time"

type Tournament struct {
	ID                string    `gorm:"primaryKey" json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Game              string    `json:"game"`
	BannerURL         string    `json:"banner_url"`
	MinRank           int       `json:"min_rank"`
	MaxRank           int       `json:"max_rank"`
	RegistrationStart time.Time `json:"registration_start"`
	RegistrationEnd   time.Time `json:"registration_end"`
	SeasonStart       time.Time `json:"season_start"`
	SeasonEnd         time.Time `json:"season_end"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Tournament) TableName() string {
	return "tournament"
}
