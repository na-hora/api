package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`

	Name        string `json:"name" gorm:"not null"`
	FantasyName string `json:"fantasyName" gorm:"not null"`
	CNPJ        string `json:"cnpj" gorm:"not null"`
	Email       string `json:"email" gorm:"not null"`
	Phone       string `json:"phone" gorm:"not null"`
	AvatarUrl   string `json:"avatarUrl"`
	CategoryID  int    `json:"categoryId" gorm:"not null"`

	CreatedAt time.Time      `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Category           CompanyCategory     `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CategoryID;references:ID"`
	CompanyHours       []CompanyHour       `json:"companyHours" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CompanyAddresses   []CompanyAddress    `json:"companyAddresses" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Users              []User              `json:"users" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CompanyPetServices []CompanyPetService `json:"services" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Appointments       []Appointment       `json:"appointments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CompanyPetHairs    []CompanyPetHair    `json:"companyPetBreeds" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CompanyPetSizes    []CompanyPetSize    `json:"companyPetSizes" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Company) TableName() string {
	return "company"
}
