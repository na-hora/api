package entity

import (
	"time"

	"gorm.io/gorm"
)

type CompanyPetServiceValue struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement:true"`

	CompanyPetServiceID uint    `json:"companyPetServiceId" gorm:"not null"`
	CompanyPetSizeID    uint    `json:"companyPetSizeId" gorm:"not null"`
	CompanyPetHairID    uint    `json:"companyPetHairId" gorm:"not null"`
	Price               float64 `json:"price" gorm:"not null"`
	ExecutionTime       int     `json:"executionTime" gorm:"not null"`

	CreatedAt time.Time      `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	CompanyPetService CompanyPetService `json:"companyPetService" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CompanyPetServiceID;references:ID;"`
	CompanyPetSize    CompanyPetSize    `json:"companyPetSize" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CompanyPetSizeID;references:ID;"`
	CompanyPetHair    CompanyPetHair    `json:"companyPetHair" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CompanyPetHairID;references:ID;"`
}

func (CompanyPetServiceValue) TableName() string {
	return "company_pet_service_value"
}
