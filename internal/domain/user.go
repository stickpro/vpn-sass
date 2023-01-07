package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Email      string         `json:"email" gorm:"unique;not null"`
	Password   string         `json:"password" gorm:"not null"`
	Status     bool           `json:"status" gorm:"not null;default:false"`
	TelegramId int            `json:"telegram_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
