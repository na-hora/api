package services

import (
	"na-hora/api/internal/entity"
	tokenDTOs "na-hora/api/internal/models/token/dtos"
	repositories "na-hora/api/internal/models/token/repositories"
	"na-hora/api/internal/utils"

	"github.com/google/uuid"
)

type TokenService interface {
	Generate(data tokenDTOs.GenerateTokenRequestBody) (*entity.Token, *utils.AppError)
	GetValidToken(key uuid.UUID) (*entity.Token, *utils.AppError)
	UseToken(key uuid.UUID, companyID uuid.UUID) *utils.AppError
}

type tokenService struct {
	tokenRepository repositories.TokenRepository
}

func GetTokenService() TokenService {
	tokenRepository := repositories.GetTokenRepository()
	return &tokenService{
		tokenRepository,
	}
}

func (ts *tokenService) Generate(data tokenDTOs.GenerateTokenRequestBody) (*entity.Token, *utils.AppError) {
	tokenCreated, err := ts.tokenRepository.Generate(data.Note)
	if err != nil {
		return nil, err
	}
	return tokenCreated, nil
}

func (ts *tokenService) GetValidToken(key uuid.UUID) (*entity.Token, *utils.AppError) {
	tokenExistent, err := ts.tokenRepository.GetByKey(key)
	if err != nil {
		return nil, err
	}
	return tokenExistent, nil
}

func (ts *tokenService) UseToken(key uuid.UUID, companyID uuid.UUID) *utils.AppError {
	err := ts.tokenRepository.MarkAsUsed(key, companyID)

	if err != nil {
		return err
	}

	return nil
}
