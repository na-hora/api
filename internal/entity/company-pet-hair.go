package entity

import (
	"time"

	"github.com/google/uuid"
)

type CompanyPetHair struct {
	ID        int       `json:"id" gorm:"primaryKey;type:integer;autoIncrement"`
	Name      string    `json:"name" gorm:"not null"`
	CompanyID uuid.UUID `json:"companyId" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`

	Company Company `json:"company" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CompanyID;references:ID;"`
}

func (CompanyPetHair) TableName() string {
	return "company_pet_hair"
}
