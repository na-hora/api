package entity

import (
	"time"
)

type PetBreed struct {
	ID        int       `json:"id" gorm:"primaryKey;type:integer;autoIncrement"`
	Name      string    `json:"name" gorm:"not null;unique"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (PetBreed) TableName() string {
	return "pet_breed"
}
