package service_test

import (
	"mercado/acoes/database"
	"os"
	"testing"

	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	dbUri := os.Getenv("TEST_DATABASE_URI")
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
