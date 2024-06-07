package services

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/company/dtos"
	repositories "na-hora/api/internal/models/company/repositories"
	"na-hora/api/internal/utils"

	"github.com/google/uuid"
)

type CompanyServiceInterface interface {
	CreateCompany(companyCreate dtos.CreateCompanyRequestBody) (*entity.Company, *utils.AppError)
	CreateAddress(companyID uuid.UUID, addressCreate dtos.CreateCompanyAddressRequestBody) (*entity.CompanyAddress, *utils.AppError)
}

type CompanyService struct {
	companyRepository repositories.CompanyRepositoryInterface
}

func GetCompanyService(repo repositories.CompanyRepositoryInterface) CompanyServiceInterface {
	return &CompanyService{
		repo,
	}
}

func (cs *CompanyService) CreateCompany(companyCreate dtos.CreateCompanyRequestBody) (*entity.Company, *utils.AppError) {
	companyCreated, err := cs.companyRepository.Create(companyCreate)
	if err != nil {
		return nil, err
	}

	return companyCreated, nil
}

func (cs *CompanyService) CreateAddress(companyID uuid.UUID, addressCreate dtos.CreateCompanyAddressRequestBody) (*entity.CompanyAddress, *utils.AppError) {
	companyAddressCreated, err := cs.companyRepository.CreateAddress(companyID, addressCreate)
	if err != nil {
		return nil, err
	}

	return companyAddressCreated, nil
}
