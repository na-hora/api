package entity

import (
	"time"

	"github.com/google/uuid"
)

type CompanyBlocklist struct {
	CompanyID uuid.UUID `json:"companyId"`
	ClientID  uuid.UUID `json:"clientId"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`

	Company Company `json:"company" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Client  Client  `json:"client" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (CompanyBlocklist) TableName() string {
	return "company_blocklist"
}
