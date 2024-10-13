package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Budget struct {
  ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Value    int64 `json:"value"`
  UserID   uuid.UUID    `json:"user_id" gorm:"type:uuid;"`
	User     User
  Currency string `json:"currency"`
	Validity int    `json:"validity"`
}

func (b *Budget) BeforeCreate(d *gorm.DB) (err error) {
  b.ID = uuid.New()
  return
}
