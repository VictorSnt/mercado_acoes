package DTO

import "time"

type CreateTransaction struct {
	UserID          uint
	EquitieID       uint
	Type            string
	Quantity        uint
	TransactionDate time.Time
}

type DisplayTransaction struct {
	ID              uint      `json:"transaction_id"`
	UserID          uint      `json:"user_id"`
	EquitieID       uint      `json:"equitie_id"`
	Type            string    `json:"type"`
	Quantity        uint      `json:"quantity"`
	UnitPrice       float64   `json:"unit_price"`
	TransactionDate time.Time `json:"transaction_date"`
}
