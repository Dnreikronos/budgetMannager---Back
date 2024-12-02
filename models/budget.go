package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Budget struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Value    int64     `json:"value"`
	UserID   uuid.UUID `json:"user_id" gorm:"type:uuid;"`
	User     User
	Currency string    `json:"currency"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
}

type BudgetInput struct {
	Value    int64     `json:"value" binding:"required,gt=0"`
	Currency string    `json:"currency" binding:"required"`
	Start    time.Time `json:"start" binding:"required"`
	End      time.Time `json:"end" binding:"required"`
}

func (b *Budget) BeforeCreate(d *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}
