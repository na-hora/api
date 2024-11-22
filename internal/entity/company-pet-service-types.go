package entity

import (
	"time"

	"gorm.io/gorm"
)

type CompanyPetServiceTypes struct {
	CompanyPetServiceID int            `json:"companyPetServiceId" gorm:"primaryKey;not null"`
	CompanyPetTypeID    int            `json:"companyPetTypeId" gorm:"primaryKey;not null"`
	CreatedAt           time.Time      `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time      `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt           gorm.DeletedAt `json:"-" gorm:"index"`

	CompanyPetService CompanyPetService `json:"companyPetService" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CompanyPetServiceID;references:ID;"`
	CompanyPetType    CompanyPetType    `json:"companyPetType" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CompanyPetTypeID;references:ID;"`
}

func (CompanyPetServiceTypes) TableName() string {
	return "company_pet_service_types"
}
