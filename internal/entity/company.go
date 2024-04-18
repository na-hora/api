package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID          uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string         `json:"name" gorm:"not null"`
	FantasyName string         `json:"fantasyName" gorm:"not null"`
	Cnpj        string         `json:"cnpj" gorm:"not null"`
	Email       string         `json:"email" gorm:"not null"`
	Phone       string         `json:"phone" gorm:"not null"`
	AvatarUrl   string         `json:"avatarUrl"`
	CategoryID  int            `json:"categoryId" gorm:"not null"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
