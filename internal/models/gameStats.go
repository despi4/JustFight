package models

// GameStats - модель для хранения статистики боя
type GameStats struct {
	ID               uint `json:"id" db:"id"`                                 // Уникальный ID записи
	MatchID          uint `json:"match_id" db:"match_id"`                     // ID матча, к которому относится статистика
	PlayerID         uint `json:"player_id" db:"player_id"`                   // ID игрока
	DamageDealt      int  `json:"damage_dealt" db:"damage_dealt"`             // Урон, нанесённый игроком
	DamageReceived   int  `json:"damage_received" db:"damage_received"`       // Урон, полученный игроком
}
