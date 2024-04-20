package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`

	CompanyID uuid.UUID `json:"companyId" gorm:"not null"`
	Username  string    `json:"username" gorm:"not null;unique"`
	Password  string    `json:"password" gorm:"not null"`

	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`

	Company Company `json:"company" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (User) TableName() string {
	return "user"
}
