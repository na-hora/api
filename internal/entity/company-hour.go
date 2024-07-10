package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyHour struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement:true"`

	CompanyID uuid.UUID `json:"companyId" gorm:"not null"`
	Weekday   int       `json:"weekday" gorm:"not null"`
	StartHour int       `json:"startHour" gorm:"not null"`
	EndHour   int       `json:"endHour" gorm:"not null"`

	CreatedAt time.Time      `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Company Company `json:"company" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CompanyID;references:ID;"`
}

func (CompanyHour) TableName() string {
	return "company_hour"
}
