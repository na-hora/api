package repositories

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/company/dtos"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyRepositoryInterface interface {
	Create(dtos.CreateCompanyRequestBody) (*entity.Company, *utils.AppError)
	CreateAddress(companyID uuid.UUID, insert dtos.CreateCompanyAddressRequestBody) (*entity.CompanyAddress, *utils.AppError)
}

type CompanyRepository struct {
	db *gorm.DB
}

func GetCompanyRepository(db *gorm.DB) CompanyRepositoryInterface {
	return &CompanyRepository{db}
}

func (cr *CompanyRepository) Create(insert dtos.CreateCompanyRequestBody) (*entity.Company, *utils.AppError) {
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

func (cr *CompanyRepository) CreateAddress(companyID uuid.UUID, insert dtos.CreateCompanyAddressRequestBody) (*entity.CompanyAddress, *utils.AppError) {
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
