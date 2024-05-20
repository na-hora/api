package services

import (
	"na-hora/api/internal/entity"
	tokenDTOs "na-hora/api/internal/models/token/dtos"
	repositories "na-hora/api/internal/models/token/repositories"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/google/uuid"
)

type TokenService interface {
	Generate(data tokenDTOs.GenerateTokenRequestBody) (*entity.Token, *utils.AppError)
	UseToken(key uuid.UUID) *utils.AppError
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

func (ts *tokenService) UseToken(key uuid.UUID) *utils.AppError {
	tokenExistent, err := ts.tokenRepository.GetByKey(key)
	if err != nil {
		return err
	}

	if tokenExistent == nil {
		return &utils.AppError{
			Message:    "Invalid validator",
			StatusCode: http.StatusUnauthorized,
		}
	}

	err = ts.tokenRepository.MarkAsUsed(key)

	if err != nil {
		return err
	}

	return nil
}
