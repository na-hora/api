package dtos

type CreateCompanyHourRequestBody struct {
	Registers []CompanyHourRequestData `json:"registers" validate:"required"`
}

type CompanyHourRequestData struct {
	ID          int `json:"id"`
	Weekday     int `json:"weekday" validate:"required"`
	StartMinute int `json:"startMinute" validate:"required"`
	EndMinute   int `json:"endMinute" validate:"required"`
}
