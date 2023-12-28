package repository

import "peken-be/models/domain"

type RoleRepository interface {
	FindByID(roleId uint) (*domain.Role, error)
	FindAll() ([]domain.Role, error)
}
