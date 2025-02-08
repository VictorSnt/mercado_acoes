package services

import (
	"encoding/json"
	"mercado/acoes/database/repositories"
	DTO "mercado/acoes/dto"
	"net/http"
)

type UsersHttpService struct {
	UserRepository repositories.UsersRepository
}

func (us UsersHttpService) CreateUser(newUser DTO.CreateUser) ([]byte, int) {
	err := us.UserRepository.Create(newUser)
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusBadRequest
	}

	response, _ := json.Marshal(map[string]string{"message": "User created successfully."})
	return response, http.StatusCreated
}

func (us UsersHttpService) FindUserById(id uint) ([]byte, int) {
	user, err := us.UserRepository.FindById(id)
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

func (us UsersHttpService) FindAllUsers() ([]byte, int) {
	users, err := us.UserRepository.FindAll()
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
