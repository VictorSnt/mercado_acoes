package database

import (
	"mercado/acoes/database/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConnection(dbUri string) (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(dbUri), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	err = Migrate(db)

	if err != nil {
		panic("failed to migrate database")
	}

	return db
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Equitie{},
		&models.Transaction{},
	)
	return err
}
