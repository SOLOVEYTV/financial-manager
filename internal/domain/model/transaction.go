package model

import "time"

type Transaction struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	Type        string
	Amount      int64
	Description *string
	Category    string
	Date        time.Time
}

func NewTransaction(
	id string,
	now time.Time,
	trType string,
	amount int64,
	description *string,
	category string,
	date time.Time,
) *Transaction {
	return &Transaction{
		ID:          id,
		CreatedAt:   now,
		UpdatedAt:   now,
		Type:        trType,
		Amount:      amount,
		Description: description,
		Category:    category,
		Date:        date,
	}
}
