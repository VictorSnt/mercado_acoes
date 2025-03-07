package repository_test

import (
	"mercado/acoes/database/repositories"
	DTO "mercado/acoes/dto"
	"testing"
	"time"

	"gorm.io/gorm"
)

func CreateTransaction(t *testing.T, tx *gorm.DB, newTransaction DTO.CreateTransaction) {
	err := repositories.TransactionsRepository{Db: tx}.Create(newTransaction)

	if err != nil {
		t.Fatalf("Error creating transaction %s: %v", newTransaction.Type, err)
	}
}

func TestCreateNewTransaction(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)
	createUser(t, tx, DTO.CreateUser{Name: "Test User", Balance: 1000})
	createEquitie(t, tx, DTO.CreateEquitie{
		Name:                  "Test Equitie",
		CurrentPrice:          1000,
		PriceChangePercentage: 10,
	})

	CreateTransaction(t, tx, DTO.CreateTransaction{
		UserID:          1,
		EquitieID:       1,
		Type:            "buy",
		Quantity:        10,
		TransactionDate: time.Now().Truncate(24 * time.Hour),
	})
}

func TestFindTransactionByUserID(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)
	createUser(t, tx, DTO.CreateUser{Name: "Test User", Balance: 1000})
	createEquitie(t, tx, DTO.CreateEquitie{
		Name:                  "Test Equitie",
		CurrentPrice:          1000,
		PriceChangePercentage: 10,
	})

	CreateTransaction(t, tx, DTO.CreateTransaction{
		UserID:          1,
		EquitieID:       1,
		Type:            "buy",
		Quantity:        10,
		TransactionDate: time.Now().Truncate(24 * time.Hour),
	})
	_, err := repositories.TransactionsRepository{Db: tx}.FindByUserId(1)

	if err != nil {
		t.Errorf("Falha ao buscar transação %s", err)
	}
}
