package entity

import (
	"time"

	"github.com/google/uuid"
)

type CompanyAddress struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CompanyID    uuid.UUID `json:"companyId" gorm:"not null"`
	ZipCode      string    `json:"zipCode"`
	CityID       uint      `json:"cityId"`
	Neighborhood string    `json:"neighborhood"`
	Street       string    `json:"street"`
	Number       string    `json:"number"`
	Complement   string    `json:"complement"`
	CreatedAt    time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`

	Company Company `json:"company" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	City    City    `json:"city" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (CompanyAddress) TableName() string {
	return "company_address"
}
