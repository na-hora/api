package dtos

import "time"

type CreateCompanyHourBlockRequestBody struct {
	Registers []CompanyHourBlockRequestData `json:"registers" validate:"required"`
}

type CompanyHourBlockRequestData struct {
	Day       time.Time `json:"day" validate:"required"`
	StartHour int       `json:"startHour" validate:"required,format=datetime"`
	EndHour   int       `json:"endHour" validate:"required,format=datetime"`
}
