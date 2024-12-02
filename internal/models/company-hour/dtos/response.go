package dtos

type ListHoursByCompanyIDResponse struct {
	ID          uint `json:"id"`
	Weekday     int  `json:"weekday"`
	StartMinute int  `json:"startMinute"`
	EndMinute   int  `json:"endMinute"`
}
