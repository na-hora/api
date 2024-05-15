package dto

import (
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/klassmann/cpfcnpj"
)

type CompanyCreate struct {
	Name        string `json:"name" gorm:"not null;unique"`
	FantasyName string `json:"fantasyName" gorm:"not null"`
	Cnpj        string `json:"cnpj" gorm:"not null;unique"`
	Email       string `json:"email" gorm:"not null"`
	Phone       string `json:"phone" gorm:"not null"`
	AvatarUrl   string `json:"avatarUrl"`
}

func (c *CompanyCreate) ValidateCNPJ() *utils.AppError {
	r := cpfcnpj.ValidateCNPJ(c.Cnpj)

	if !r {
		return &utils.AppError{
			Message:    "invalid CNPJ",
			StatusCode: http.StatusBadRequest,
		}
	}

	return nil
}
