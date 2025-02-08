package handlers

import (
	"encoding/json"
	"mercado/acoes/database/models"
	"mercado/acoes/database/repositories"
	DTO "mercado/acoes/dto"
	"net/http"

	"gorm.io/gorm"
)

func CreateEquitie(db *gorm.DB, newEquitie DTO.CreateEquitie) ([]byte, int) {
	EquitiesRepository := repositories.EquitiesRepository{Db: db}
	err := EquitiesRepository.Create(newEquitie)
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusBadRequest
	}

	response, _ := json.Marshal(map[string]string{"message": "Equitie created successfully."})
	return response, http.StatusCreated
}

func FindEquitieById(db *gorm.DB, id uint) ([]byte, int) {
	EquitiesRepository := repositories.EquitiesRepository{Db: db}
	equitie, err := EquitiesRepository.FindById(id)
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusNotFound
	}
	response, err := json.Marshal(equitie)

	if err != nil {
		errorResponse := map[string]string{
			"error":  err.Error(),
			"detail": "Error while parsing equitie to json.",
		}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusInternalServerError
	}

	return response, http.StatusOK
}

func FindAllEquities(db *gorm.DB) ([]byte, int) {
	EquitiesRepository := repositories.EquitiesRepository{Db: db}
	equities, err := EquitiesRepository.FindAll()
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusNotFound
	}

	response, err := json.Marshal(equities)
	if err != nil {
		errorResponse := map[string]string{
			"error":  err.Error(),
			"detail": "Error while parsing equities to json.",
		}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusInternalServerError
	}
	return response, http.StatusOK
}

func UpdateEquitieName(db *gorm.DB, id uint, updatedEquitie DTO.UpdateEquitie) ([]byte, int) {
	EquitiesRepository := repositories.EquitiesRepository{Db: db}
	err := EquitiesRepository.Update(id, models.Equitie{Name: updatedEquitie.Name})
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusNotFound
	}

	response, _ := json.Marshal(map[string]string{"message": "Equitie updated successfully."})
	return response, http.StatusOK
}

func UpdateEquitiePrice(db *gorm.DB, id uint, updatedEquitie DTO.UpdateEquitiePrice) ([]byte, int) {
	EquitiesRepository := repositories.EquitiesRepository{Db: db}
	err := EquitiesRepository.Update(id, updatedEquitie)
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusNotFound
	}

	response, _ := json.Marshal(map[string]string{"message": "Equitie updated successfully."})
	return response, http.StatusOK
}
