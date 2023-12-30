package repository

import "peken-be/models/domain"

type MemberRepository interface {
	FindByID(userId uint) (*domain.Member, error)
	FindAll() ([]domain.Member, error)
	Save(member *domain.Member) (*domain.Member, error)
	Update(member *domain.Member) (*domain.Member, error)
	Delete(member *domain.Member) error
}
