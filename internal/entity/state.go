package entity

import (
	"time"
)

type State struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	UF        string    `json:"uf" gorm:"not null;unique"`
	Name      string    `json:"name" gorm:"not null;unique"`
	IBGE      int       `json:"ibge" gorm:"not null"`
	DDD       string    `json:"ddd"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
}
