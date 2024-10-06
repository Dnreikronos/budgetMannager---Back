package models

type Budget struct {
	ID       int64  `json:"id"`
	Value    string `json:"value"`
	UserID   User    `json:"user_id"`
	Currency string `json:"currency"`
	Validity int    `json:"validity"`
}
