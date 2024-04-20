package entity

import (
	"time"
)

type City struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name      string    `json:"name" gorm:"not null"`
	StateID   uint      `json:"stateId" gorm:"not null"`
	IBGE      string    `json:"ibge" gorm:"not null"`
	LatLon    string    `json:"latLon"`
	CodTom    uint      `json:"codTom"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`

	State State `json:"state" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (City) TableName() string {
	return "city"
}
