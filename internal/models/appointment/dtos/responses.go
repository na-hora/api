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

	StartTime   time.Time `json:"startTime"`
	TotalTime   int       `json:"totalTime"`
	ServiceName string    `json:"serviceName"`
}

type Appointment struct {
	ID          uuid.UUID `json:"id"`
	ServiceName string    `json:"serviceName"`
	StartTime   string    `json:"startTime"`
	TotalTime   int       `json:"totalTime"`
	TotalPrice  float64   `json:"totalPrice"`
	Canceled    bool      `json:"canceled"`
}
