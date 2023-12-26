package repository

import (
	"peken-be/models/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

// NewUserRepository returns new UserRepositoryImpl.
func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	db.AutoMigrate(&domain.User{})
	return &UserRepositoryImpl{
		DB: db,
	}
}

// Save ...
func (repository *UserRepositoryImpl) Save(user *domain.User) (*domain.User, error) {
	error := repository.DB.Create(&user).Error
	if error != nil {
		return nil, error
	}
	return user, error
}

// Update ...
func (repository *UserRepositoryImpl) Update(user *domain.User) (*domain.User, error) {
	error := repository.DB.Save(&user).Error
	if error != nil {
		return nil, error
	}
	return user, error
}

// Delete ...
func (repository *UserRepositoryImpl) Delete(user *domain.User) error {
	error := repository.DB.Delete(user).Error
	return error
}

// FindByID ...
func (repository *UserRepositoryImpl) FindByID(userId uint) (*domain.User, error) {
	user := domain.User{}
	error := repository.DB.First(&user, userId).Error
	if error != nil {
		return nil, error
	}
	return &user, error
}

// FindAll ...
func (repository *UserRepositoryImpl) FindAll() ([]domain.User, error) {
	var users []domain.User
	error := repository.DB.Find(&users).Error
	return users, error
}

// FindByUsername ...
func (repository *UserRepositoryImpl) FindByUsername(username string) (*domain.User, error) {
	user := domain.User{}
	error := repository.DB.Where("username = ?", username).First(&user).Error
	if error != nil {
		return nil, error
	}
	return &user, error
}
