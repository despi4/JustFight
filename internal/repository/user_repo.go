package repository

import "just-fight/internal/models"

type UserRepository interface {
	Create(user *models.User) error
	GetById(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

func (u *models.User) Create(user *models.User) error {
	return nil
}
