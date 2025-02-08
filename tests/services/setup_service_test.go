package service_test

import (
	"mercado/acoes/database"
	"mercado/acoes/database/repositories"
	DTO "mercado/acoes/dto"
	"os"
	"testing"

	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	dbUri := "../test.db"
	db = database.GetConnection(dbUri)
	code := m.Run()
	os.Exit(code)
}

func SetupTest(t *testing.T) (*gorm.DB, func(t *testing.T)) {
	tx := db.Begin()
	if tx.Error != nil {
		t.Fatalf("Failed to begin transaction: %v", tx.Error)
	}
	return tx, func(t *testing.T) {
		tx.Rollback()
	}
}

func CreateUser(t *testing.T, tx *gorm.DB, newUser DTO.CreateUser) {
	err := repositories.UsersRepository{Db: tx}.Create(newUser)

	if err != nil {
		t.Fatalf("Error creating user %s: %v", newUser.Name, err)
	}
}

func CreateEquitie(t *testing.T, tx *gorm.DB, newEquitie DTO.CreateEquitie) {
	err := repositories.EquitiesRepository{Db: tx}.Create(newEquitie)

	if err != nil {
		t.Fatalf("Error creating equitie %s: %v", newEquitie.Name, err)
	}
}
