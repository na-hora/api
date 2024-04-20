package entity

import (
	"time"
)

type PetSize struct {
	ID        int       `json:"id" gorm:"primaryKey;type:integer;autoIncrement"`
	Name      string    `json:"name" gorm:"not null;unique"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (PetSize) TableName() string {
	return "pet_size"
}
