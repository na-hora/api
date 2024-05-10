package entity

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid();"`
	Name      string    `json:"name" gorm:"not null"`
	Phone     string    `json:"phone" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`

	Appointments []Appointment `json:"appointments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Client) TableName() string {
	return "client"
}
