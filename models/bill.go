package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Bills struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Value     int64     `json:"value" gorm:"not null"`
	BudgetID  uuid.UUID `json:"budget_id" gorm:"type:uuid;not null"`
	Budget    Budget    `json:"-" gorm:"foreignKey:BudgetID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Category  string    `json:"category" gorm:"not null"`
	Status    string    `json:"status" gorm:"default:'pending'"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type BillInput struct {
	Value    int64  `json:"value" binding:"required,gt=0"`
	Category string `json:"category" binding:"required"`
	Status   string `json:"status" binding:"required,oneof='paid' 'unpaid' 'pending'"`
}

func (b *Bills) BeforeCreate(d *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}
