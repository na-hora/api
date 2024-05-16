package usecases

import (
	"na-hora/api/internal/dto"
	repositories "na-hora/api/internal/models/company/repositories"
	"na-hora/api/internal/utils"
)

type CompanyUsecase interface {
	CreateCompany(companyCreate dto.CompanyCreate) *utils.AppError
}

type companyUsecase struct {
	companyRepository repositories.CompanyRepository
}

func GetCompanyUsecase() CompanyUsecase {
	companyRepo := repositories.GetCompanyRepository()
	return &companyUsecase{
		companyRepo,
	}
}

func (uc *companyUsecase) CreateCompany(companyCreate dto.CompanyCreate) *utils.AppError {
	err := uc.companyRepository.Create(companyCreate)
	if err != nil {
		return err
	}

	return nil
}
