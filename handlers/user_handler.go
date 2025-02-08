package handlers

import (
	"encoding/json"
	"mercado/acoes/database/repositories"
	DTO "mercado/acoes/dto"
	"net/http"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, newUser DTO.CreateUser) ([]byte, int) {
	UserRepository := repositories.UsersRepository{Db: db}
	err := UserRepository.Create(newUser)
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusBadRequest
	}

	response, _ := json.Marshal(map[string]string{"message": "User created successfully."})
	return response, http.StatusCreated
}

func FindUserById(db *gorm.DB, id uint) ([]byte, int) {
	UserRepository := repositories.UsersRepository{Db: db}
	user, err := UserRepository.FindById(id)
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusNotFound
	}
	response, err := json.Marshal(user)

	if err != nil {
		errorResponse := map[string]string{
			"error":  err.Error(),
			"detail": "Error while parsing user to json.",
		}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusInternalServerError
	}

	return response, http.StatusOK
}

func FindAllUsers(db *gorm.DB) ([]byte, int) {
	UserRepository := repositories.UsersRepository{Db: db}
	users, err := UserRepository.FindAll()
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusNotFound
	}

	response, err := json.Marshal(users)
	if err != nil {
		errorResponse := map[string]string{
			"error":  err.Error(),
			"detail": "Error while parsing users to json.",
		}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusInternalServerError
	}
	return response, http.StatusOK
}

func UpdateUser(db *gorm.DB, id uint, updatedUser DTO.UpdateUser) ([]byte, int) {
	UserRepository := repositories.UsersRepository{Db: db}
	err := UserRepository.Update(id, updatedUser)
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusBadRequest
	}

	response, _ := json.Marshal(map[string]string{"message": "User updated successfully."})
	return response, http.StatusOK
}
