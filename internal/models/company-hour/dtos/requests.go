package dtos

type CreateCompanyHourRequestBody struct {
	Registers []CompanyHourRequestData `json:"registers" validate:"required"`
}

type CompanyHourRequestData struct {
	Weekday   int `json:"weekday" validate:"required"`
	StartHour int `json:"startHour" validate:"required,format=datetime"`
	EndHour   int `json:"endHour" validate:"required,format=datetime"`
}
