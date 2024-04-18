package entity

import (
	"time"
)

type City struct {
	ID        int       `json:"id" gorm:"primaryKey;type:integer;autoIncrement"`
	Name      string    `json:"name" gorm:"not null"`
	StateID   int       `json:"stateId" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`

	State State `json:"state" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
