package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID          uint
	EquitieID       uint
	Type            string
	Quantity        uint
	UnitPrice       float64
	TransactionDate time.Time `gorm:"type:date"`
}
