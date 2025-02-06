package DTO

import "time"

type CreateTransaction struct {
	UserID          uint
	EquitieID       uint
	Type            string
	Quantity        uint
	UnitPrice       float64
	TransactionDate time.Time
}
