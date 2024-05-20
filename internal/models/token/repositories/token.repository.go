package repositories

import (
	config "na-hora/api/configs"
	"na-hora/api/internal/entity"
	"na-hora/api/internal/utils"
	"net/http"

	"gorm.io/gorm"
)

type TokenRepository interface {
	Generate(note string) (*entity.Token, *utils.AppError)
}

type tokenRepository struct {
	db *gorm.DB
}

func GetTokenRepository() TokenRepository {
	db := config.DB
	return &tokenRepository{db}
}

func (t *tokenRepository) Generate(note string) (*entity.Token, *utils.AppError) {
	insertValue := entity.Token{
		Note: note,
	}

	data := t.db.Create(&insertValue)
	if data.Error != nil {
		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &insertValue, nil
}
