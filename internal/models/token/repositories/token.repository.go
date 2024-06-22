package repositories

import (
	"na-hora/api/internal/entity"
	"na-hora/api/internal/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TokenRepositoryInterface interface {
	Generate(note string) (*entity.Token, *utils.AppError)
	GetValidByKey(key uuid.UUID) (*entity.Token, *utils.AppError)
	MarkAsUsedByCompany(key uuid.UUID, companyID uuid.UUID, tx *gorm.DB) *utils.AppError
	MarkAsUsedByUser(key uuid.UUID, userID uuid.UUID, tx *gorm.DB) *utils.AppError
}

type TokenRepository struct {
	db *gorm.DB
}

func GetTokenRepository(db *gorm.DB) TokenRepositoryInterface {
	return &TokenRepository{db}
}

func (t *TokenRepository) Generate(note string) (*entity.Token, *utils.AppError) {
	now := time.Now()
	afterOneDay := now.Add(time.Hour * 24)

	insertValue := entity.Token{
		Note:      note,
		ExpiresAt: &afterOneDay,
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

func (t *TokenRepository) GetValidByKey(key uuid.UUID) (*entity.Token, *utils.AppError) {
	var token entity.Token
	data := t.db.Where("key = ? and used = false and expires_at > now()", key).First(&token)
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

func (t *TokenRepository) MarkAsUsedByCompany(key uuid.UUID, companyID uuid.UUID, tx *gorm.DB) *utils.AppError {
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

func (t *TokenRepository) MarkAsUsedByUser(key uuid.UUID, userID uuid.UUID, tx *gorm.DB) *utils.AppError {
	if tx == nil {
		tx = t.db
	}

	data := tx.Model(&entity.Token{}).Where("key = ?", key).Update("used", true).Update("user_id", userID)
	if data.Error != nil {
		return &utils.AppError{
			Message:    data.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return nil
}
