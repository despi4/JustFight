package repository

import "just-fight/internal/models"

type MatchRepository interface {
	Create(match *models.Match) error
	GetByID(id uint) (*models.Match, error)
	GetByPlayerID(PlayerID uint) ([]*models.Match, error)
	GetRecentMatches(limit int) ([]*models.Match, error)
}
