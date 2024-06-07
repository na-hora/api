//go:build wireinject
// +build wireinject

package injector

import (
	companyRepositories "na-hora/api/internal/models/company/repositories"
	companyServices "na-hora/api/internal/models/company/services"

	userRepositories "na-hora/api/internal/models/user/repositories"
	userServices "na-hora/api/internal/models/user/services"

	tokenRepositories "na-hora/api/internal/models/token/repositories"
	tokenServices "na-hora/api/internal/models/token/services"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var CompanyServiceSet = wire.NewSet(
	companyRepositories.GetCompanyRepository,
	companyServices.GetCompanyService,
)

func InitializeCompanyService(db *gorm.DB) companyServices.CompanyServiceInterface {
	wire.Build(CompanyServiceSet)
	return nil // This line should never be executed; Wire replaces it
}

// ------------------------------------------------------------------------ //

var UserServiceSet = wire.NewSet(
	userRepositories.GetUserRepository,
	userServices.GetUserService,
)

func InitializeUserService(db *gorm.DB) userServices.UserServiceInterface {
	wire.Build(UserServiceSet)
	return nil // This line should never be executed; Wire replaces it
}

// ------------------------------------------------------------------------ //

var TokenServiceSet = wire.NewSet(
	tokenRepositories.GetTokenRepository,
	tokenServices.GetTokenService,
)

func InitializeTokenService(db *gorm.DB) tokenServices.TokenServiceInterface {
	wire.Build(TokenServiceSet)
	return nil // This line should never be executed; Wire replaces it
}
