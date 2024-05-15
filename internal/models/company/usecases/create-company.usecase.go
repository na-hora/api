package companyUsecase

import (
	"na-hora/api/internal/dto"
	repositories "na-hora/api/internal/models/company/repositories"
	"na-hora/api/internal/utils"
)

type CompanyUsecase struct {
	companyRepository repositories.CompanyRepository
}

func NewCompanyUsecase(companyRepo repositories.CompanyRepository) *CompanyUsecase {
	return &CompanyUsecase{
		companyRepo,
	}
}

func (uc *CompanyUsecase) CreateCompany(companyCreate dto.CompanyCreate) *utils.AppError {
	err := uc.companyRepository.Create(companyCreate)
	if err != nil {
		return err
	}

	return nil
}
