package entity

import (
	"time"

	"github.com/google/uuid"
)

type CompanyPetBreed struct {
	ID         int       `json:"id" gorm:"primaryKey;type:integer;autoIncrement"`
	CompanyID  uuid.UUID `json:"companyId" gorm:"not null"`
	PetBreedID int       `json:"petBreedId" gorm:"not null"`
	ExtraValue float64   `json:"extraValue"`
	ExtraTime  float64   `json:"extraTime"`
	CreatedAt  time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`

	Company  Company  `json:"company" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PetBreed PetBreed `json:"petBreed" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (CompanyPetBreed) TableName() string {
	return "company_pet_breed"
}
