package models

import "time"

type Bills struct {
	ID        int64     `json:"id"`
	Value     string    `json:"value"`
	UserID    int       `json:"userid"`
	BudgetID  int       `json:"budgetid"`
	Category  string    `json:"category"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
