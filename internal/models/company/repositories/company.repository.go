package repositories

import (
	config "na-hora/api/configs"
	"na-hora/api/internal/dto"
	"na-hora/api/internal/entity"
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

func GetCompanyRepository() CompanyRepository {
	db := config.DB
	return &companyRepository{db}
}

func (cr *companyRepository) Create(insert dto.CompanyCreate) *utils.AppError {
	insertValue := entity.Company{
		CNPJ:        insert.Cnpj,
		Name:        insert.Name,
		FantasyName: insert.FantasyName,
		Email:       insert.Email,
		Phone:       insert.Phone,
		AvatarUrl:   insert.AvatarUrl,
		CategoryID:  1, // TODO: enum
	}

	data := cr.db.Create(&insertValue)

	if data.Error != nil {
		return &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil
}
