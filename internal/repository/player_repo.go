package repository

import "just-fight/internal/models"

type PlayerRepository interface {
	// Basic CRUD methods
	Create(player *models.Player) error
	GetById(id uint) (*models.Player, error)
	Update(player *models.Player) error
	Delete(id uint) error
	List() ([]*models.Player, error)

	// Specific method for search
	GetByName(name string) (*models.Player, error)
	FindByMinHP(minHP int) ([]*models.Player, error)
	FindByAttackRange(minAttack, maxAttack int) ([]*models.Player, error)

	// Game methods
	UpdateHP(id uint, newHP int) error
	IncreaseHP(id uint, amount int) error
	DecreaseHP(id uint, amount int) error
	UpdateStats(id uint, attack, speed int) error

	// Batch-Operation
	BulkUpdateHP(updates map[uint]int) error
	GetTopPlayersByStat(stat string, limit int) ([]*models.Player, error)
}
