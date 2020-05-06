package models

import "github.com/lib/pq"

type Teams struct {
	ID      uint           `json:"id" gorm:"primary_key"`
	Name    string         `json:"name" gorm:"unique;not null"`
	Players pq.StringArray `json:"players" gorm:"unique;not null;type:varchar(100)[]"`
}
