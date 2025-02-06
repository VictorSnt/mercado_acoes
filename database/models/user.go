package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Name    string
	Balance float64
}
