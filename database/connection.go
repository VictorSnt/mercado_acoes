package database

import (
	"mercado/acoes/database/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConnection() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("database/project.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Migrate(db)
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Equitie{},
		&models.Usuario{},
		&models.Transaction{},
	)
}
