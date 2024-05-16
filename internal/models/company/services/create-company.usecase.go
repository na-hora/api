package services

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/models/company/dtos"
	repositories "na-hora/api/internal/models/company/repositories"
	"na-hora/api/internal/utils"
)

type CompanyService interface {
	CreateCompany(companyCreate dtos.CreateCompanyRequestBody) (*entity.Company, *utils.AppError)
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

func (uc *companyService) CreateCompany(companyCreate dtos.CreateCompanyRequestBody) (*entity.Company, *utils.AppError) {
	companyCreated, err := uc.companyRepository.Create(companyCreate)
	if err != nil {
		return nil, err
	}

	return companyCreated, nil
}
