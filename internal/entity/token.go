package entity

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	Key       uuid.UUID  `json:"key" gorm:"primaryKey;default:gen_random_uuid();type:uuid"`
	Note      string     `json:"note"`
	CompanyID *uuid.UUID `json:"companyId"`
	UserID    *uuid.UUID `json:"userId"`
	Used      bool       `json:"used" gorm:"not null;default:false"`

	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`

	ExpiresAt *time.Time `json:"expiresAt"`

	Company Company `json:"company" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CompanyID;references:ID;"`
	User    User    `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UserID;references:ID;"`
}

func (Token) TableName() string {
	return "token"
}
