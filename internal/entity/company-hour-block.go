package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyHourBlock struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement:true"`

	CompanyID uuid.UUID `json:"companyId" gorm:"not null"`
	Day       time.Time `json:"day" gorm:"not null"`
	StartHour time.Time `json:"startHour" gorm:"not null"`
	EndHour   time.Time `json:"endHour" gorm:"not null"`

	CreatedAt time.Time      `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Company Company `json:"company" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CompanyID;references:ID;"`
}

func (CompanyHourBlock) TableName() string {
	return "company_hour_block"
}
