package repository_test

import (
	"mercado/acoes/database/repositories"
	DTO "mercado/acoes/dto"
	"mercado/acoes/handlers"
	"net/http"
	"testing"

	"gorm.io/gorm"
)

func createUser(t *testing.T, tx *gorm.DB, newUser DTO.CreateUser) {
	err := repositories.UsersRepository{Db: tx}.Create(newUser)

	if err != nil {
		t.Fatalf("Error creating user %s: %v", newUser.Name, err)
	}
}

func TestCreateNewUser(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)

	createUser(t, tx, DTO.CreateUser{Name: "Test User", Balance: 1000})
}

func TestGetUserById(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)

	createUser(t, tx, DTO.CreateUser{Name: "Test User", Balance: 1000})

	user, err := repositories.UsersRepository{Db: tx}.FindById(1)

	if err != nil {
		t.Errorf("Error getting user by id: %v", err)
	}

	if user.ID != 1 {
		t.Errorf("User id is not 1")
	}
}

func TestFindListOfUsers(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)

	createUser(t, tx, DTO.CreateUser{Name: "Test User", Balance: 1000})
	createUser(t, tx, DTO.CreateUser{Name: "Test User2", Balance: 2000})
	createUser(t, tx, DTO.CreateUser{Name: "Test User3", Balance: 3000})

	users, err := repositories.UsersRepository{Db: tx}.FindAll()

	if err != nil {
		t.Errorf("Error getting list of users: %v", err)
	}

	if len(users) != 3 {
		t.Errorf("Users list is not 3, is %v", len(users))
	}
}

func TestDeleteUser(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)

	createUser(t, tx, DTO.CreateUser{Name: "Test User", Balance: 1000})

	err := repositories.UsersRepository{Db: tx}.Delete(1)

	if err != nil {
		t.Errorf("Error deleting user: %v", err)
	}

	_, err = repositories.UsersRepository{Db: tx}.FindAll()

	if err == nil {
		t.Errorf("User was not deleted")
	}

}

func TestFindUserEquities(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)

	createUser(t, tx, DTO.CreateUser{Name: "Test User", Balance: 1000})
	createEquitie(t, tx, DTO.CreateEquitie{Name: "Test Equitie", CurrentPrice: 100, PriceChangePercentage: 0.1})

	newTransaction := DTO.CreateTransaction{
		UserID:    1,
		EquitieID: 1,
		Type:      "BUY",
		Quantity:  10,
	}

	response, status := handlers.CreateEquiteTransaction(tx, newTransaction)

	if status != http.StatusCreated {
		t.Fatalf("Error creating transaction %v", response)
	}

	userStocks, err := repositories.UsersRepository{Db: tx}.FindEquities(1)

	if err != nil {
		t.Fatalf("Error getting user equities %v", err)
	}

	if len(userStocks) != 1 {
		t.Errorf("User equities list is not 1, is %v", len(userStocks))
	}

	if userStocks[0].TotalQuantity != 10 {
		t.Errorf("User equities total quantity is not 10, is %v", userStocks[0].TotalQuantity)
	}
}
