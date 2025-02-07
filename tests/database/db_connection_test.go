package test

import (
	"mercado/acoes/database"
	"os"
	"testing"
)

func TestGetConnectionAndMigrations(t *testing.T) {
	dbUri := os.Getenv("TEST_DATABASE_URI")
	db := database.GetConnection(dbUri)

	result := db.Raw("SELECT * FROM users").Scan(&struct{}{})

	if result.Error != nil {
		t.Errorf("Migrations fail: %v", result.Error)
	}
}
