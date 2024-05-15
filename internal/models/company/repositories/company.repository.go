package repositories

import (
	"na-hora/api/internal/dto"
	"na-hora/api/internal/utils"
	"net/http"

	"gorm.io/gorm"
)

type CompanyRepository interface {
	Create(dto.CompanyCreate) *utils.AppError
}

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyRepository{db: db}
}

func (c *companyRepository) Create(company dto.CompanyCreate) *utils.AppError {
	data := c.db.Create(&company)
	if data.Error != nil {
		return &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil
}
