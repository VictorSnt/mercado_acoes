package service_test

import (
	"mercado/acoes/database/repositories"
	DTO "mercado/acoes/dto"
	"mercado/acoes/enums"
	"mercado/acoes/handlers"
	"net/http"
	"testing"
	"time"
)

func TestCreateEquiteTransaction(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)

	CreateUser(t, tx, DTO.CreateUser{Name: "User 1", Balance: 100})
	CreateEquitie(t, tx, DTO.CreateEquitie{Name: "Equitie 1", CurrentPrice: 10})

	response, status := handlers.CreateEquiteTransaction(tx, DTO.CreateTransaction{
		UserID:          1,
		EquitieID:       1,
		Quantity:        10,
		Type:            string(enums.TransactionBuyOperation),
		TransactionDate: time.Now().Truncate(24 * time.Hour),
	})

	if status != http.StatusCreated {
		t.Log(response)
		t.Fatalf("Error creating transaction: %v", status)
	}
}

func TestBuyEquiteWithInsufficientBalance(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)

	CreateUser(t, tx, DTO.CreateUser{Name: "User 1", Balance: 0})
	CreateEquitie(t, tx, DTO.CreateEquitie{Name: "Equitie 1", CurrentPrice: 10})

	response, status := handlers.CreateEquiteTransaction(tx, DTO.CreateTransaction{
		UserID:          1,
		EquitieID:       1,
		Quantity:        10,
		Type:            string(enums.TransactionBuyOperation),
		TransactionDate: time.Now().Truncate(24 * time.Hour),
	})

	if status != http.StatusBadRequest {
		t.Log(response)
		t.Fatalf("Error creating transaction: %v", status)
	}

	validResponse := map[string]string{
		"detail": "invalid transaction for userID 1 .",
		"error":  "insufficient balance",
	}
	if response["error"] != validResponse["error"] {
		t.Fatalf("Error creating transaction, wrong error response: %v", response)
	}
}

func TestSellEquiteWithInsufficientQuantity(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)

	CreateUser(t, tx, DTO.CreateUser{Name: "User 1", Balance: 100})
	CreateEquitie(t, tx, DTO.CreateEquitie{Name: "Equitie 1", CurrentPrice: 10})

	response, status := handlers.CreateEquiteTransaction(tx, DTO.CreateTransaction{
		UserID:          1,
		EquitieID:       1,
		Quantity:        10,
		Type:            string(enums.TransactionSaleOperation),
		TransactionDate: time.Now().Truncate(24 * time.Hour),
	})

	if status != http.StatusBadRequest {
		t.Log(response)
		t.Fatalf("Error creating transaction: %v", status)
	}

	validResponse := map[string]string{
		"detail": "invalid transaction for userID 1 .",
		"error":  "insufficient equities for sale operation",
	}

	if response["error"] != validResponse["error"] {
		t.Fatalf("Error creating transaction, wrong error response: %v", response)
	}
}

func TestEquitePriceIncrement(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)

	CreateUser(t, tx, DTO.CreateUser{Name: "User 1", Balance: 150})
	CreateEquitie(t, tx, DTO.CreateEquitie{Name: "Equitie 1", CurrentPrice: 10, PriceChangePercentage: 0.1})

	handlers.CreateEquiteTransaction(tx, DTO.CreateTransaction{
		UserID:          1,
		EquitieID:       1,
		Quantity:        5,
		Type:            string(enums.TransactionBuyOperation),
		TransactionDate: time.Now().Truncate(24 * time.Hour),
	})

	response, status := handlers.CreateEquiteTransaction(tx, DTO.CreateTransaction{
		UserID:          1,
		EquitieID:       1,
		Quantity:        5,
		Type:            string(enums.TransactionBuyOperation),
		TransactionDate: time.Now().Truncate(24 * time.Hour),
	})

	equitieDTO, err := repositories.EquitiesRepository{Db: tx}.FindById(1)
	if err != nil {
		t.Fatalf("Error finding equitie: %v", err)
	}

	if equitieDTO.CurrentPrice == 10 {
		t.Fatalf("Error equite price was not incremented after first sale: %v", equitieDTO.CurrentPrice)
	}

	if status != http.StatusCreated {
		t.Log(response)
		t.Fatalf("Error equite prece was not incremented after first sale: %v", status)
	}
}

func TestEquitePriceDecrement(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)

	CreateUser(t, tx, DTO.CreateUser{Name: "User 1", Balance: 100})
	CreateEquitie(t, tx, DTO.CreateEquitie{Name: "Equitie 1", CurrentPrice: 10.0, PriceChangePercentage: 3})

	handlers.CreateEquiteTransaction(tx, DTO.CreateTransaction{
		UserID:          1,
		EquitieID:       1,
		Quantity:        5,
		Type:            string(enums.TransactionBuyOperation),
		TransactionDate: time.Now().Truncate(24 * time.Hour),
	})

	handlers.CreateEquiteTransaction(tx, DTO.CreateTransaction{
		UserID:          1,
		EquitieID:       1,
		Quantity:        4,
		Type:            string(enums.TransactionBuyOperation),
		TransactionDate: time.Now().Truncate(24 * time.Hour),
	})

	response, status := handlers.CreateEquiteTransaction(tx, DTO.CreateTransaction{
		UserID:          1,
		EquitieID:       1,
		Quantity:        5,
		Type:            string(enums.TransactionSaleOperation),
		TransactionDate: time.Now().Truncate(24 * time.Hour),
	})

	equitieDTO, err := repositories.EquitiesRepository{Db: tx}.FindById(1)
	if err != nil {
		t.Fatalf("Error finding equitie: %v", err)
	}

	if equitieDTO.CurrentPrice == 10 {
		t.Fatalf("Error equite price was not decremented after first sale: %v", equitieDTO.CurrentPrice)
	}

	if status != http.StatusCreated {
		t.Log(response)
		t.Fatalf("Error equite prece was not decremented after first sale: %v", status)
	}
}
