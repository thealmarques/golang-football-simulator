package models

type Player struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name" gorm:"not null"`
	Age         int    `json:"age" gorm:"not null"`
	GoalKeeping int    `json:"goalkeeping"`
	Attack      int    `json:"attack"`
	Defense     int    `json:"defense"`
	Creativity  int    `json:"creativity"`
}
