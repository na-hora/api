package entity

import (
	"time"

	"github.com/google/uuid"
)

type CompanyPetSize struct {
	ID         int       `json:"id" gorm:"primaryKey;type:integer;autoIncrement"`
	CompanyID  uuid.UUID `json:"companyId" gorm:"not null"`
	PetSizeID  int       `json:"petSizeId" gorm:"not null"`
	ExtraValue float64   `json:"extraValue" gorm:"not null"`
	ExtraTime  float64   `json:"extraTime" gorm:"not null"`
	CreatedAt  time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`

	Company Company `json:"company" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PetSize PetSize `json:"petSize" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (CompanyPetSize) TableName() string {
	return "company_pet_size"
}
