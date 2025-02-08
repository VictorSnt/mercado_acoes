package service_test

import (
	DTO "mercado/acoes/dto"
	"mercado/acoes/handlers"
	"testing"
)

func TestCreateUserServiceCreateUser(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)
	response, status := handlers.CreateUser(tx, DTO.CreateUser{
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

	response, status := handlers.FindUserById(tx, 1)
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

	response, status := handlers.FindAllUsers(tx)
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

	_, status := handlers.CreateUser(tx, DTO.CreateUser{
		Name:    "John Doe",
		Balance: 1000,
	})

	if status != 201 {
		t.Errorf("Expected status 201, got %d", status)
	}

	response, status := handlers.FindUserById(tx, 1)

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

	_, status := handlers.CreateUser(tx, DTO.CreateUser{
		Name:    "John Doe",
		Balance: 1000,
	})

	if status != 201 {
		t.Errorf("Expected status 201, got %d", status)
	}

	response, status := handlers.FindAllUsers(tx)

	if status != 200 {
		t.Errorf("Expected status 200, got %d", status)
	}

	expectedResponse := `[{"user_id":1,"name":"John Doe","balance":1000}]`
	if string(response) != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, string(response))
	}
}

func TestUpdateUser(t *testing.T) {
	tx, teardown := SetupTest(t)
	defer teardown(t)

	_, status := handlers.CreateUser(tx, DTO.CreateUser{
		Name:    "John Doe",
		Balance: 1000,
	})

	if status != 201 {
		t.Errorf("Expected status 201, got %d", status)
	}

	response, status := handlers.UpdateUser(tx, 1, DTO.UpdateUser{
		Name: "Jane Doe",
	})

	if status != 200 {
		t.Errorf("Expected status 200, got %d", status)
	}

	expectedResponse := `{"message":"User updated successfully."}`
	if string(response) != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, string(response))
	}
}
