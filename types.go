package main

import "golang.org/x/exp/rand"

type User struct {
	ID       int64
	Email    string
	Name     string
	Password string
}

func newUser(Name, Email, Password string) *User {
	return &User{
		ID          int64  `json:"id"`
		Email:    Email,
		Name:     Name,
		Password: Password,
	}
}

type Budget struct {
	ID    int
	Value string
}

func newBudget(Valeu string) *Budget {
	ID:       rand.Intn(10000),
	Value: value,
}

tpye Bills struct {
	ID int
	Value string
	Data 
}