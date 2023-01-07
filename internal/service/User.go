package service

import (
	"github.com/stickpro/vpn-sass/internal/domain"
	"github.com/stickpro/vpn-sass/internal/repository"
)

type UsersService struct {
	repository repository.Users
}

type UserServiceInterface interface {
	LoadAll() ([]domain.User, error)
	FindByTgId(int) (domain.User, error)
}

func NewUsersService(repository repository.Users) *UsersService {
	return &UsersService{repository: repository}
}

func (u *UsersService) LoadAll() ([]domain.User, error) {
	return u.repository.GetAll()
}

func (u *UsersService) FindByTgId(telegramId int) (domain.User, error) {
	return u.repository.FindByColumn(telegramId, "telegram_id")
}
