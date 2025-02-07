package test

import (
	"mercado/acoes/database/repositories"
	DTO "mercado/acoes/dto"
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
	tx, teardown := setupTest(t)
	defer teardown(t)

	createUser(t, tx, DTO.CreateUser{Name: "Test User", Balance: 1000})
}

func TestGetUserById(t *testing.T) {
	tx, teardown := setupTest(t)
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
	tx, teardown := setupTest(t)
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
