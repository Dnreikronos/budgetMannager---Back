package models

import "time"

type Bills struct {
	ID        int64     `json:"id"`
	Value     string    `json:"value"`
	UserID    User      `json:"user_id"`
	BudgetID  int       `json:"budget_id"`
	Category  string    `json:"category"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
