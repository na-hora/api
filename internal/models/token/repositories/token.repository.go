package repositories

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TokenRepositoryInterface interface {
	Generate(note string) (*entity.Token, *utils.AppError)
	GetByKey(key uuid.UUID) (*entity.Token, *utils.AppError)
	MarkAsUsed(key uuid.UUID, companyID uuid.UUID, tx *gorm.DB) *utils.AppError
}

type TokenRepository struct {
	db *gorm.DB
}

func GetTokenRepository(db *gorm.DB) TokenRepositoryInterface {
	return &TokenRepository{db}
}

func (t *TokenRepository) Generate(note string) (*entity.Token, *utils.AppError) {
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

func (t *TokenRepository) GetByKey(key uuid.UUID) (*entity.Token, *utils.AppError) {
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

func (t *TokenRepository) MarkAsUsed(key uuid.UUID, companyID uuid.UUID, tx *gorm.DB) *utils.AppError {
	if tx == nil {
		tx = t.db
	}

	data := tx.Model(&entity.Token{}).Where("key = ?", key).Update("used", true).Update("company_id", companyID)
	if data.Error != nil {
		return &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return nil
}
