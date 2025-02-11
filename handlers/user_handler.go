package handlers

import (
	"mercado/acoes/database/repositories"
	DTO "mercado/acoes/dto"
	"net/http"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, newUser DTO.CreateUser) (map[string]string, int) {
	UserRepository := repositories.UsersRepository{Db: db}
	err := UserRepository.Create(newUser)
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		return errorResponse, http.StatusBadRequest
	}

	response := map[string]string{"message": "User created successfully."}
	return response, http.StatusCreated
}

func FindUserById(db *gorm.DB, id uint) (interface{}, int) {
	UserRepository := repositories.UsersRepository{Db: db}
	user, err := UserRepository.FindById(id)
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		return errorResponse, http.StatusNotFound
	}

	return user, http.StatusOK
}

func FindAllUsers(db *gorm.DB) (interface{}, int) {
	UserRepository := repositories.UsersRepository{Db: db}
	users, err := UserRepository.FindAll()

	if err != nil {
		errorResponse := map[string]string{
			"error":  err.Error(),
			"detail": "Error while seaching user.",
		}
		return errorResponse, http.StatusNotFound
	}

	return users, http.StatusOK
}

func UpdateUser(db *gorm.DB, id uint, updatedUser DTO.UpdateUser) (map[string]string, int) {
	UserRepository := repositories.UsersRepository{Db: db}
	err := UserRepository.Update(id, updatedUser)
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		return errorResponse, http.StatusBadRequest
	}

	response := map[string]string{"message": "User updated successfully."}
	return response, http.StatusOK
}

func DeleteUser(db *gorm.DB, id uint) (map[string]string, int) {
	UserRepository := repositories.UsersRepository{Db: db}
	err := UserRepository.Delete(id)
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		return errorResponse, http.StatusBadRequest
	}

	response := map[string]string{"message": "User deleted successfully."}
	return response, http.StatusOK
}

func RetriveUserEquitieStocks(db *gorm.DB, userID uint) (interface{}, int) {
	UserRepository := repositories.UsersRepository{Db: db}
	userStock, err := UserRepository.FindEquities(userID)
	if err != nil {
		errorResponse := map[string]string{
			"error":  err.Error(),
			"detail": "User not found or user has no equities.",
		}
		return errorResponse, http.StatusNotFound
	}

	return userStock, http.StatusOK
}
