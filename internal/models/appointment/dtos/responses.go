package dtos

import (
	"time"

	"github.com/google/uuid"
)

type ListAppointmentsResponse struct {
	Appointments []Appointment `json:"appointments"`
}

type CreateAppointmentResponse struct {
	ID uuid.UUID `json:"id"`

	StartTime time.Time `json:"startTime"`
	TotalTime int       `json:"totalTime"`
}

type Appointment struct {
	ID                uuid.UUID `json:"id"`
	PetName           string    `json:"petName"`
	StartTime         string    `json:"startTime"`
	TotalTime         int       `json:"totalTime"`
	TotalPrice        float64   `json:"totalPrice"`
	PaymentMode       string    `json:"paymentMode"`
	Canceled          bool      `json:"canceled"`
	CancelationReason string    `json:"cancelationReason"`
}
