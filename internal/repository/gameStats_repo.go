package repository

import "just-fight/internal/models"

type GameStatsRepository interface {
	GetByPlayerID(playerID uint) (*models.GameStats, error)
	UpdateStats(playerID uint, wins, losses int, rating float64) error
	IncrementWins(playerID uint) error
	IncrementLosses(playerID uint) error
	UpdateRating(playerID uint, newRating float64) error
	GetTopPlayers(limit int) ([]*models.GameStats, error)
}
