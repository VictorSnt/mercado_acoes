package handlers

import (
	"mercado/acoes/database/models"
	"mercado/acoes/database/repositories"
	DTO "mercado/acoes/dto"
	"net/http"

	"gorm.io/gorm"
)

func CreateEquitie(db *gorm.DB, newEquitie DTO.CreateEquitie) (interface{}, int) {
	EquitiesRepository := repositories.EquitiesRepository{Db: db}
	err := EquitiesRepository.Create(newEquitie)
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		return errorResponse, http.StatusBadRequest
	}

	successResponse := map[string]string{"message": "Equitie created successfully."}
	return successResponse, http.StatusCreated
}

func FindEquitieById(db *gorm.DB, id uint) (interface{}, int) {
	EquitiesRepository := repositories.EquitiesRepository{Db: db}
	equitie, err := EquitiesRepository.FindById(id)
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		return errorResponse, http.StatusNotFound
	}
	return equitie, http.StatusOK
}

func FindAllEquities(db *gorm.DB) (interface{}, int) {
	EquitiesRepository := repositories.EquitiesRepository{Db: db}
	equities, err := EquitiesRepository.FindAll()
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		return errorResponse, http.StatusNotFound
	}
	return equities, http.StatusOK
}

func UpdateEquitieName(db *gorm.DB, id uint, updatedEquitie DTO.UpdateEquitie) (interface{}, int) {
	EquitiesRepository := repositories.EquitiesRepository{Db: db}
	err := EquitiesRepository.Update(id, models.Equitie{Name: updatedEquitie.Name})
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		return errorResponse, http.StatusNotFound
	}

	successResponse := map[string]string{"message": "Equitie updated successfully."}
	return successResponse, http.StatusOK
}
