package dtos

import (
	"time"

	"github.com/google/uuid"
)

type CreateAppointmentParams struct {
	ClientID                 uuid.UUID `json:"clientId"`
	CompanyPetServiceValueID int       `json:"companyPetServiceValueId"`
	StartTime                time.Time `json:"startTime"`
	PetName                  string    `json:"petName"`
	PaymentMode              string    `json:"paymentMode"`
	Note                     string    `json:"note"`
	TotalTime                int       `json:"totalTime"`
	TotalPrice               float64   `json:"totalPrice"`
	Canceled                 bool      `json:"canceled"`
}
