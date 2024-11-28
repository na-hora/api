package dtos

import "github.com/google/uuid"

type CreateCompanyHourParams struct {
	Weekday     int
	StartMinute int
	EndMinute   int
	CompanyID   uuid.UUID
}
