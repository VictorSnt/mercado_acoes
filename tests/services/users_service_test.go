package service_test

import (
	"mercado/acoes/database/repositories"
	DTO "mercado/acoes/dto"
	"mercado/acoes/services"
	"testing"
)

func TestCreateUserServiceCreateUser(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)

	userHttpService := services.UsersHttpService{
		UserRepository: repositories.UsersRepository{Db: tx},
	}

	response, status := userHttpService.CreateUser(DTO.CreateUser{
		Name:    "John Doe",
		Balance: 1000,
	})

	if status != 201 {
		t.Errorf("Expected status 201, got %d", status)
	}

	expectedResponse := `{"message":"User created successfully."}`
	if string(response) != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, string(response))
	}
}

func TestFailToFindUserById(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)

	userHttpService := services.UsersHttpService{
		UserRepository: repositories.UsersRepository{Db: tx},
	}

	response, status := userHttpService.FindUserById(1)
	if status != 404 {
		t.Errorf("Expected status 404, got %d", status)
	}

	expectedResponse := `{"error":"record not found"}`
	if string(response) != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, string(response))
	}
}

func TestFailToFindAllUsers(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)

	userHttpService := services.UsersHttpService{
		UserRepository: repositories.UsersRepository{Db: tx},
	}

	response, status := userHttpService.FindAllUsers()
	if status != 404 {
		t.Errorf("Expected status 404, got %d", status)
	}

	expectedResponse := `{"error":"record not found"}`
	if string(response) != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, string(response))
	}
}

func TestFindUserById(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)

	userHttpService := services.UsersHttpService{
		UserRepository: repositories.UsersRepository{Db: tx},
	}

	_, status := userHttpService.CreateUser(DTO.CreateUser{
		Name:    "John Doe",
		Balance: 1000,
	})

	if status != 201 {
		t.Errorf("Expected status 201, got %d", status)
	}

	response, status := userHttpService.FindUserById(1)

	if status != 200 {
		t.Errorf("Expected status 200, got %d", status)
	}

	expectedResponse := `{"user_id":1,"name":"John Doe","balance":1000}`
	if string(response) != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, string(response))
	}
}

func TestFindAllUsers(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)

	userHttpService := services.UsersHttpService{
		UserRepository: repositories.UsersRepository{Db: tx},
	}

	_, status := userHttpService.CreateUser(DTO.CreateUser{
		Name:    "John Doe",
		Balance: 1000,
	})

	if status != 201 {
		t.Errorf("Expected status 201, got %d", status)
	}

	response, status := userHttpService.FindAllUsers()

	if status != 200 {
		t.Errorf("Expected status 200, got %d", status)
	}

	expectedResponse := `[{"user_id":1,"name":"John Doe","balance":1000}]`
	if string(response) != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, string(response))
	}
}
