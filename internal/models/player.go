package models

type Player struct {
	ID     uint   `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	HP     int    `json:"hp" db:"hp"`
	Attack int    `json:"attack" db:"attack"`
	Speed  int    `json:"speed" db:"speed"`
}
