package repositories

import (
	config "na-hora/api/configs"
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/company/dtos"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyRepository interface {
	Create(dtos.CreateCompanyRequestBody) (*entity.Company, *utils.AppError)
	CreateAddress(companyID uuid.UUID, insert dtos.CreateCompanyAddressRequestBody) (*entity.CompanyAddress, *utils.AppError)
}

type companyRepository struct {
	db *gorm.DB
}

func GetCompanyRepository() CompanyRepository {
	db := config.DB
	return &companyRepository{db}
}

func (cr *companyRepository) Create(insert dtos.CreateCompanyRequestBody) (*entity.Company, *utils.AppError) {
	insertValue := entity.Company{
		CNPJ:        insert.CNPJ,
		Name:        insert.Name,
		FantasyName: insert.FantasyName,
		Email:       insert.Email,
		Phone:       insert.Phone,
		AvatarUrl:   insert.AvatarUrl,
		CategoryID:  1, // TODO: enum
	}

	data := cr.db.Create(&insertValue)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &insertValue, nil
}

func (cr *companyRepository) CreateAddress(companyID uuid.UUID, insert dtos.CreateCompanyAddressRequestBody) (*entity.CompanyAddress, *utils.AppError) {
	insertValue := entity.CompanyAddress{
		CompanyID:    companyID,
		Neighborhood: insert.Neighborhood,
		Street:       insert.Street,
		Number:       insert.Number,
		Complement:   insert.Complement,
		ZipCode:      insert.ZipCode,
		CityID:       insert.CityID,
	}

	data := cr.db.Create(&insertValue)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &insertValue, nil
}
