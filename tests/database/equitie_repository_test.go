package repository_test

import (
	"mercado/acoes/database/repositories"
	DTO "mercado/acoes/dto"
	"testing"

	"gorm.io/gorm"
)

func createEquitie(t *testing.T, tx *gorm.DB, newEquitie DTO.CreateEquitie) {
	err := repositories.EquitiesRepository{Db: tx}.Create(newEquitie)

	if err != nil {
		t.Fatalf("Error creating equitie %s: %v", newEquitie.Name, err)
	}
}

func TestCreateNewEquitie(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)

	createEquitie(t, tx, DTO.CreateEquitie{
		Name:                  "Test Equitie",
		CurrentPrince:         1000,
		PriceChangePercentage: 10,
	})
}

func TestGetEquitieById(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)

	createEquitie(t, tx, DTO.CreateEquitie{
		Name:                  "Test Equitie",
		CurrentPrince:         1000,
		PriceChangePercentage: 10,
	})

	equitie, err := repositories.EquitiesRepository{Db: tx}.FindById(1)

	if err != nil {
		t.Errorf("Error getting equitie by id: %v", err)
	}

	if equitie.ID != 1 {
		t.Errorf("Equitie id is not 1")
	}
}

func TestFindListOfEquities(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)

	createEquitie(t, tx, DTO.CreateEquitie{
		Name:                  "Test Equitie",
		CurrentPrince:         1000,
		PriceChangePercentage: 10,
	})
	createEquitie(t, tx, DTO.CreateEquitie{
		Name:                  "Test Equitie2",
		CurrentPrince:         2000,
		PriceChangePercentage: 20,
	})
	createEquitie(t, tx, DTO.CreateEquitie{
		Name:                  "Test Equitie3",
		CurrentPrince:         3000,
		PriceChangePercentage: 30,
	})

	equities, err := repositories.EquitiesRepository{Db: tx}.FindAll()

	if err != nil {
		t.Errorf("Error getting list of equities: %v", err)
	}

	if len(equities) != 3 {
		t.Errorf("Equities list is not 3, is %v", len(equities))
	}
}
