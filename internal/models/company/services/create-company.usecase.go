package services

import (
	"na-hora/api/internal/dto"
	repositories "na-hora/api/internal/models/company/repositories"
	"na-hora/api/internal/utils"
)

type CompanyService interface {
	CreateCompany(companyCreate dto.CompanyCreate) *utils.AppError
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

func (uc *companyService) CreateCompany(companyCreate dto.CompanyCreate) *utils.AppError {
	err := uc.companyRepository.Create(companyCreate)
	if err != nil {
		return err
	}

	return nil
}
