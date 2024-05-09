package entity

import (
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	ID                       uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid();"`
	CompanyID                uuid.UUID `json:"companyId" gorm:"not null"`
	ClientID                 uuid.UUID `json:"clientId" gorm:"not null"`
	CompanyPetServiceValueID uuid.UUID `json:"CompanyPetServiceValueId" gorm:"not null"`
	PetName                  string    `json:"petName"`
	StartTime                time.Time `json:"startTime" gorm:"not null"`
	TotalTime                int       `json:"totalTime" gorm:"not null"`
	TotalPrice               float64   `json:"totalPrice" gorm:"not null"`
	PaymentMode              string    `json:"paymentMode"`
	Canceled                 bool      `json:"canceled"`
	CancelationReason        string    `json:"cancelationReason"`
	CreatedAt                time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt                time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`

	Company                Company                `json:"company" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CompanyID;references:ID;"`
	Client                 Client                 `json:"client" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:ClientID;references:ID;"`
	CompanyPetServiceValue CompanyPetServiceValue `json:"companyPetServiceValue" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CompanyPetServiceValueID;references:ID;"`
}

func (Appointment) TableName() string {
	return "appointment"
}
