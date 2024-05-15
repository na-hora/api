//go:build wireinject
// +build wireinject

package main

import (
	"na-hora/api/internal/models/company/repositories"
	companyUsecase "na-hora/api/internal/models/company/usecases"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func NewCompanyUseCase(db *gorm.DB) *companyUsecase.CompanyUsecase {
	wire.Build(
		repositories.NewCompanyRepository,
		companyUsecase.NewCompanyUsecase,
	)

	return &companyUsecase.CompanyUsecase{}
}
