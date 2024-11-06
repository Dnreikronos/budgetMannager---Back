package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Bills struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Value     int64     `json:"value"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;"`
	User      User
	BudgetID  uuid.UUID `json:"budget_id" gorm:"type:uuid;"`
	Budget    Budget
	Category  string    `json:"category"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BillInput struct {
	Value    int64  `json:"value"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

func (b *Bills) BeforeCreate(d *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}
