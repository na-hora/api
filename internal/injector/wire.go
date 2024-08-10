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

	stateRepositories "na-hora/api/internal/models/state/repositories"
	stateServices "na-hora/api/internal/models/state/services"

	cityRepositories "na-hora/api/internal/models/city/repositories"
	cityServices "na-hora/api/internal/models/city/services"

	companyHourRepositories "na-hora/api/internal/models/company-hour/repositories"
	companyHourServices "na-hora/api/internal/models/company-hour/services"

	companyHourBlockRepositories "na-hora/api/internal/models/company-hour-block/repositories"
	companyHourBlockServices "na-hora/api/internal/models/company-hour-block/services"

	petServiceRepositories "na-hora/api/internal/models/pet-service/repositories"
	petServiceServices "na-hora/api/internal/models/pet-service/services"

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

// ------------------------------------------------------------------------ //

var StateServiceSet = wire.NewSet(
	stateRepositories.GetStateRepository,
	stateServices.GetStateService,
)

func InitializeStateService(db *gorm.DB) stateServices.StateServiceInterface {
	wire.Build(StateServiceSet)
	return nil // This line should never be executed; Wire replaces it
}

// ------------------------------------------------------------------------ //

var CityServiceSet = wire.NewSet(
	cityRepositories.GetCityRepository,
	cityServices.GetCityService,
)

func InitializeCityService(db *gorm.DB) cityServices.CityServiceInterface {
	wire.Build(CityServiceSet)
	return nil // This line should never be executed; Wire replaces it
}

// ------------------------------------------------------------------------ //

var CompanyHourServiceSet = wire.NewSet(
	companyHourRepositories.GetCompanyHourRepository,
	companyHourServices.GetCompanyHourService,
)

func InitializeCompanyHourService(db *gorm.DB) companyHourServices.CompanyHourServiceInterface {
	wire.Build(CompanyHourServiceSet)
	return nil // This line should never be executed; Wire replaces it
}

// ------------------------------------------------------------------------ //

var CompanyHourBlockServiceSet = wire.NewSet(
	companyHourBlockRepositories.GetCompanyHourBlockRepository,
	companyHourBlockServices.GetCompanyHourBlockService,
)

func InitializeCompanyHourBlockService(db *gorm.DB) companyHourBlockServices.CompanyHourBlockServiceInterface {
	wire.Build(CompanyHourBlockServiceSet)
	return nil // This line should never be executed; Wire replaces it
}

// ------------------------------------------------------------------------ //

var PetServiceServiceSet = wire.NewSet(
	petServiceRepositories.GetPetServiceRepository,
	petServiceServices.GetPetServiceService,
)

func InitializePetServiceService(db *gorm.DB) petServiceServices.PetServiceInterface {
	wire.Build(PetServiceServiceSet)
	return nil // This line should never be executed; Wire replaces it
}
