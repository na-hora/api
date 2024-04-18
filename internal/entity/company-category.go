package entity

import (
	"time"

	"gorm.io/gorm"
)

type CompanyCategory struct {
	ID        int            `json:"id" gorm:"primaryKey;type:integer;autoIncrement"`
	Name      string         `json:"name" gorm:"not null"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
