package services

import (
	"na-hora/api/internal/entity"
	tokenDTOs "na-hora/api/internal/models/token/dtos"
	repositories "na-hora/api/internal/models/token/repositories"
	"na-hora/api/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TokenServiceInterface interface {
	Generate(data tokenDTOs.GenerateTokenRequestBody) (*entity.Token, *utils.AppError)
	GetValidToken(key uuid.UUID) (*entity.Token, *utils.AppError)
	UseCompanyToken(key uuid.UUID, companyID uuid.UUID, tx *gorm.DB) *utils.AppError
	UseUserToken(key uuid.UUID, userID uuid.UUID, tx *gorm.DB) *utils.AppError
}

type TokenService struct {
	tokenRepository repositories.TokenRepositoryInterface
}

func GetTokenService(repo repositories.TokenRepositoryInterface) TokenServiceInterface {
	return &TokenService{
		repo,
	}
}

func (ts *TokenService) Generate(data tokenDTOs.GenerateTokenRequestBody) (*entity.Token, *utils.AppError) {
	tokenCreated, err := ts.tokenRepository.Generate(data.Note)
	if err != nil {
		return nil, err
	}
	return tokenCreated, nil
}

func (ts *TokenService) GetValidToken(key uuid.UUID) (*entity.Token, *utils.AppError) {
	tokenExistent, err := ts.tokenRepository.GetValidByKey(key)
	if err != nil {
		return nil, err
	}
	return tokenExistent, nil
}

func (ts *TokenService) UseCompanyToken(key uuid.UUID, companyID uuid.UUID, tx *gorm.DB) *utils.AppError {
	err := ts.tokenRepository.MarkAsUsedByCompany(key, companyID, tx)

	if err != nil {
		return err
	}

	return nil
}

func (ts *TokenService) UseUserToken(key uuid.UUID, userID uuid.UUID, tx *gorm.DB) *utils.AppError {
	err := ts.tokenRepository.MarkAsUsedByUser(key, userID, tx)

	if err != nil {
		return err
	}

	return nil
}
