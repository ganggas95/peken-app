package repository

import (
	"peken-be/models/domain"

	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	DB *gorm.DB
}

// NewRoleRepository returns new RoleRepository.
func NewRoleRepository(db *gorm.DB) *RoleRepositoryImpl {
	db.AutoMigrate(&domain.Role{})
	return &RoleRepositoryImpl{
		DB: db,
	}
}

// FindByID ...
func (repository *RoleRepositoryImpl) FindByID(roleId uint) (*domain.Role, error) {
	role := domain.Role{}
	err := repository.DB.First(&role, roleId).Error
	return &role, err
}

// FindAll ...
func (repository *RoleRepositoryImpl) FindAll() ([]domain.Role, error) {
	var roles []domain.Role
	err := repository.DB.Find(&roles).Error
	return roles, err
}
