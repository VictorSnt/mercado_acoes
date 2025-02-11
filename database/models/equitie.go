package models

import "gorm.io/gorm"

type Equitie struct {
	gorm.Model
	Name                  string `gorm:"unique"`
	CurrentPrice          float64
	PriceChangePercentage float64
}
