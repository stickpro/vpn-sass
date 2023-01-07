package repository

import (
	"github.com/stickpro/vpn-sass/internal/domain"
	"github.com/stickpro/vpn-sass/pkg/logger"
	"gorm.io/gorm"
)

type Users interface {
	Create(domain.User) (domain.User, error)
	GetAll() ([]domain.User, error)
	FindByColumn(any, string) (domain.User, error)
	Migrate() error
}

type userRepository struct {
	DB *gorm.DB
}

func newUserRepository(DB *gorm.DB) *userRepository {
	return &userRepository{DB: DB}
}

func (u userRepository) Migrate() error {
	logger.Info("[UserRepository]...Migrate")
	return u.DB.AutoMigrate(&domain.User{})
}

func (u userRepository) Create(user domain.User) (domain.User, error) {
	logger.Info("[UserRepository]...Save")
	err := u.DB.Create(&user).Error
	return user, err
}

func (u userRepository) GetAll() (users []domain.User, err error) {
	logger.Info("[UserRepository]...Get All")
	err = u.DB.Find(&users).Error
	return users, err
}

func (u userRepository) FindByColumn(value any, columnName string) (domain.User, error) {
	logger.Info("[UserRepository]... Find by column")
	var user domain.User
	err := u.DB.Find(&user, columnName+" = ?", value).Error
	return user, err
}
