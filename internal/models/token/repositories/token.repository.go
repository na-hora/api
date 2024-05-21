package repositories

import (
	config "na-hora/api/configs"
	"na-hora/api/internal/entity"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TokenRepository interface {
	Generate(note string) (*entity.Token, *utils.AppError)
	GetByKey(key uuid.UUID) (*entity.Token, *utils.AppError)
	MarkAsUsed(key uuid.UUID, companyID uuid.UUID) *utils.AppError
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

func (t *tokenRepository) GetByKey(key uuid.UUID) (*entity.Token, *utils.AppError) {
	var token entity.Token
	data := t.db.Where("key = ? and used = false", key).First(&token)
	if data.Error != nil {
		if data.Error == gorm.ErrRecordNotFound {
			return nil, &utils.AppError{
				Message:    "Invalid validator",
				StatusCode: http.StatusUnauthorized,
			}
		}

		return nil, &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return &token, nil
}

func (t *tokenRepository) MarkAsUsed(key uuid.UUID, companyID uuid.UUID) *utils.AppError {
	data := t.db.Model(&entity.Token{}).Where("key = ?", key).Update("used", true).Update("company_id", companyID)
	if data.Error != nil {
		return &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return nil
}
