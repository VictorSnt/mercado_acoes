package database

import (
	"fmt"
	"mercado/acoes/database/models"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConnection(dbUri string) *gorm.DB {
	var once sync.Once
	var db *gorm.DB

	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open(dbUri), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		fmt.Println("Banco de dados conectado!")

		err = Migrate(db)
		if err != nil {
			panic("failed to migrate database")
		}
	})

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
