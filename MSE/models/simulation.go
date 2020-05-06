package models

import "github.com/lib/pq"

type Simulation struct {
	ID      uint           `json:"id" gorm:"primary_key"`
	MatchId uint           `json:"match_id" gorm:"not null"`
	Result  string         `json:"result" gorm:"not null"`
	Events  pq.StringArray `json:"events" gorm:"not null;type:varchar(300)[]"`
}
