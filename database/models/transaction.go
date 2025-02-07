package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID          uint
	User            User
	EquitieID       uint
	Equitie         Equitie
	Type            string
	Quantity        uint
	UnitPrice       float64
	TransactionDate time.Time `gorm:"type:date"`
}
