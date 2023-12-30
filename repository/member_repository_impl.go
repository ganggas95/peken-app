package repository

import (
	"peken-be/models/domain"

	"gorm.io/gorm"
)

type MemberRepositoryImpl struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *MemberRepositoryImpl {
	db.AutoMigrate(&domain.Member{})
	return &MemberRepositoryImpl{
		db: db,
	}
}

func (repository *MemberRepositoryImpl) Save(member *domain.Member) (*domain.Member, error) {
	error := repository.db.Create(&member).Error
	if error != nil {
		return nil, error
	}
	return member, error
}

func (repository *MemberRepositoryImpl) Update(member *domain.Member) (*domain.Member, error) {
	error := repository.db.Save(&member).Error
	if error != nil {
		return nil, error
	}
	return member, error
}

func (repository *MemberRepositoryImpl) Delete(member *domain.Member) error {
	error := repository.db.Delete(&member).Error
	return error
}

func (repository *MemberRepositoryImpl) FindByID(userId uint) (*domain.Member, error) {
	member := domain.Member{}
	error := repository.db.First(&member, userId).Error
	return &member, error
}

func (repository *MemberRepositoryImpl) FindAll() ([]domain.Member, error) {
	var members []domain.Member
	error := repository.db.Find(&members).Error
	return members, error
}
