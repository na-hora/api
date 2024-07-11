package dtos

import (
	"time"

	"github.com/google/uuid"
)

type CreateCompanyHourBlockParams struct {
	Day       time.Time
	StartHour int
	EndHour   int
	CompanyID uuid.UUID
}
