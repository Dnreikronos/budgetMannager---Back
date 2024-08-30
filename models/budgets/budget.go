package models

type Budget struct {
	ID       int64  `json:"id"`
	Value    string `json:"value"`
	UserID   int    `json:"userid"`
	Currency string `json:"currency"`
	Validity int    `json:"validity"`
}
