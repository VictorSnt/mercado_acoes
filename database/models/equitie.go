package models

import "gorm.io/gorm"

type Equitie struct {
	gorm.Model
	Name                  string `gorm:"unique"`
	CurrentPrince         float64
	PriceChangePercentage float64
}
