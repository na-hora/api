package dtos

import "github.com/google/uuid"

type CreateCompanyHourParams struct {
	Weekday   int
	StartHour int
	EndHour   int
	CompanyID uuid.UUID
}
