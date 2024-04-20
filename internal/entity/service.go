package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// table service {
//   id int pk
//   company_id uuid [not null]
//   name string [not null]
//   price float [not null]
//   execution_time int [not null]
//   concurrency int [not null, default: 1]
//   created_at timestamp [default: 'now()']
//   updated_at timestamp [default: 'now()']
//   deleted_at timestamp [default: 'null']
// }
// Ref: service.company_id>company.id

type Service struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement:true"`

	CompanyID     uuid.UUID `json:"companyId" gorm:"not null"`
	Name          string    `json:"name" gorm:"not null"`
	Price         float64   `json:"price" gorm:"not null"`
	ExecutionTime int       `json:"executionTime" gorm:"not null"`
	Concurrency   int       `json:"concurrency" gorm:"not null;default:1"`

	CreatedAt time.Time      `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Company Company `json:"company" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Service) TableName() string {
	return "service"
}
