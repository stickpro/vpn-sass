package repository

import (
	"github.com/stickpro/vpn-sass/pkg/logger"
	"gorm.io/gorm"
)

type Repositories struct {
	Users Users
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users: newUserRepository(db),
	}
}

func (r Repositories) Migrate() {
	logger.Info("[Database Migration start]")
	_ = r.Users.Migrate()
}
