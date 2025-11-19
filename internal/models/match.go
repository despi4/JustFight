package models

import "time"

// Match - структура для представления боя
type Match struct {
	ID        uint      `json:"id" db:"id"`                 // Уникальный ID матча
	Player1ID uint      `json:"player1_id" db:"player1_id"` // ID первого игрока
	Player2ID uint      `json:"player2_id" db:"player2_id"` // ID второго игрока
	WinnerID  uint      `json:"winner_id" db:"winner_id"`   // ID победителя
	StartedAt time.Time `json:"started_at" db:"started_at"` // Время начала матча
	EndedAt   time.Time `json:"ended_at" db:"ended_at"`     // Время завершения матча
	Duration  int       `json:"duration" db:"duration"`     // Длительность матча в секундах
	Result    string    `json:"result" db:"result"`         // Результат матча (например, победа игрока 1 или 2)
}
