// models/matches.go

package models

type Matches struct {
	ID     uint `json:"id" gorm:"primary_key"`
	HomeId uint `json:"home_team_id"`
	AwayId uint `json:"away_team_id"`
}
