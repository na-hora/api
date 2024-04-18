package entity

import (
	"time"
)

type State struct {
	ID        int       `json:"id" gorm:"primaryKey;type:integer;autoIncrement"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
}
