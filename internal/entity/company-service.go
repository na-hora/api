package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyService struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement:true"`

	CompanyID     uuid.UUID `json:"companyId" gorm:"not null"`
	Name          string    `json:"name" gorm:"not null"`
	Price         float64   `json:"price" gorm:"not null"`
	ExecutionTime int       `json:"executionTime" gorm:"not null"`
	Concurrency   int       `json:"concurrency" gorm:"not null;default:1"`

	CreatedAt time.Time      `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Company Company `json:"company" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CompanyID;references:ID;"`
}

func (CompanyService) TableName() string {
	return "company_service"
}
