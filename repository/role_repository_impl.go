package repository

import (
	"peken-be/models/domain"

	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	DB        *gorm.DB
	BaseQuery *gorm.DB
}

// NewRoleRepository returns new RoleRepository.
func NewRoleRepository(db *gorm.DB) *RoleRepositoryImpl {
	db.AutoMigrate(&domain.Role{})
	return &RoleRepositoryImpl{
		DB:        db,
		BaseQuery: db.Model(&domain.Role{}).Preload("Users"),
	}
}

// FindByID ...
func (repository *RoleRepositoryImpl) FindByID(roleId uint) (*domain.Role, error) {
	role := domain.Role{}
	err := repository.BaseQuery.First(&role, roleId).Error
	return &role, err
}

// FindAll ...
func (repository *RoleRepositoryImpl) FindAll() ([]domain.Role, error) {
	var roles []domain.Role
	err := repository.BaseQuery.Find(&roles).Error
	return roles, err
}
