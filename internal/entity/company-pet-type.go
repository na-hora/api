package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyPetType struct {
	ID        int            `json:"id" gorm:"primaryKey;type:integer;autoIncrement"`
	Name      string         `json:"name" gorm:"not null"`
	CompanyID uuid.UUID      `json:"companyId" gorm:"not null"`
	CreatedAt time.Time      `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Company      Company                  `json:"company" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CompanyID;references:ID;"`
	ServiceTypes []CompanyPetServiceTypes `json:"serviceTypes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CompanyPetTypeID;references:ID;"`
}

func (CompanyPetType) TableName() string {
	return "company_pet_type"
}
