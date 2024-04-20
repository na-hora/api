package entity

import (
	"time"

	"gorm.io/gorm"
)

type CompanyCategory struct {
	ID int `json:"id" gorm:"primaryKey;type:integer;autoIncrement"`

	Name string `json:"name" gorm:"not null"`

	CreatedAt time.Time      `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (CompanyCategory) TableName() string {
	return "company_category"
}
