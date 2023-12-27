package repository

import (
	"peken-be/models/domain"
)

type UserRepository interface {
	Save(user *domain.User) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	Delete(user *domain.User) error
	FindByID(userId uint) (*domain.User, error)
	FindAll() ([]domain.User, error)
	FindByUsername(username string) (*domain.User, error)
}
