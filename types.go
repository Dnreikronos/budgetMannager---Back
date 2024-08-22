package main

import (
	"time"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
}

type Budget struct {
	ID       int    `json:"id"`
	Value    string `json:"value"`
	UserID   int    `json:"userid"`
	Currency string `json:"currency"`
	Validity int    `json:"validity"`
}

type Bills struct {
	ID        int       `json:"id"`
	Value     string    `json:"value"`
	UserID    int       `json:"userid"`
	BudgetID  int       `json:"budgetid"`
	Category  string    `json:"category"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// func newUser(Name, Email, Password string) *User {
// 	return &User{
// 		ID:       int64,
// 		Email:    Email,
// 		Name:     Name,
// 		Password: Password,
// 	}
// }
