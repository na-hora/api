package services

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/company/dtos"
	repositories "na-hora/api/internal/models/company/repositories"
	"na-hora/api/internal/utils"

	"github.com/google/uuid"
)

type CompanyService interface {
	CreateCompany(companyCreate dtos.CreateCompanyRequestBody) (*entity.Company, *utils.AppError)
	CreateAddress(companyID uuid.UUID, addressCreate dtos.CreateCompanyAddressRequestBody) (*entity.CompanyAddress, *utils.AppError)
}

type companyService struct {
	companyRepository repositories.CompanyRepository
}

func GetCompanyService() CompanyService {
	companyRepo := repositories.GetCompanyRepository()
	return &companyService{
		companyRepo,
	}
}

func (cs *companyService) CreateCompany(companyCreate dtos.CreateCompanyRequestBody) (*entity.Company, *utils.AppError) {
	companyCreated, err := cs.companyRepository.Create(companyCreate)
	if err != nil {
		return nil, err
	}

	return companyCreated, nil
}

func (cs *companyService) CreateAddress(companyID uuid.UUID, addressCreate dtos.CreateCompanyAddressRequestBody) (*entity.CompanyAddress, *utils.AppError) {
	companyAddressCreated, err := cs.companyRepository.CreateAddress(companyID, addressCreate)
	if err != nil {
		return nil, err
	}

	return companyAddressCreated, nil
}
