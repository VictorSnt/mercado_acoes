package handlers

import (
	"encoding/json"
	"fmt"
	"mercado/acoes/database/repositories"
	DTO "mercado/acoes/dto"
	"mercado/acoes/enums"
	"mercado/acoes/services"
	"net/http"

	"gorm.io/gorm"
)

func CreateEquiteTransaction(
	db *gorm.DB,
	newTransaction DTO.CreateTransaction,

) ([]byte, int) {

	TransactionRepository := repositories.TransactionsRepository{Db: db}
	EquitiesRepository := repositories.EquitiesRepository{Db: db}

	err := services.ValidadeUserEquitieTransaction(db, newTransaction)

	if err != nil {
		errorResponse := map[string]string{
			"error": err.Error(),
			"detail": fmt.Sprintf(
				"invalid transaction for userID %d .",
				newTransaction.UserID,
			),
		}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusBadRequest
	}
	equitieDTO, err := EquitiesRepository.FindById(newTransaction.EquitieID)
	if err != nil {
		errorResponse := map[string]string{
			"error": err.Error(),
			"detail": fmt.Sprintf(
				"error while finding equitie %d .",
				newTransaction.EquitieID,
			),
		}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusBadRequest
	}

	var transactionValue float64 = float64(newTransaction.Quantity) * equitieDTO.CurrentPrince
	err = services.UpdateUserBalance(
		db, newTransaction.UserID,
		transactionValue,
		newTransaction.Type,
	)
	if err != nil {
		errorResponse := map[string]string{
			"error": err.Error(),
			"detail": fmt.Sprintf(
				"error while charging user %d .",
				newTransaction.UserID,
			),
		}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusBadRequest
	}

	err = TransactionRepository.Create(newTransaction)
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusBadRequest
	}

	equiteUpdateDTO, err := services.UpdateEquitiePrice(
		db,
		newTransaction.EquitieID,
		enums.TransactionType(newTransaction.Type),
	)
	if err != nil {
		errorResponse := map[string]string{
			"error":  err.Error(),
			"detail": "Error while updating equitie price.",
		}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusBadRequest
	}

	err = EquitiesRepository.Update(newTransaction.EquitieID, equiteUpdateDTO)
	if err != nil {
		errorResponse := map[string]string{
			"error":  err.Error(),
			"detail": "Error while updating equitie price.",
		}
		response, _ := json.Marshal(errorResponse)
		return response, http.StatusBadRequest
	}

	response, _ := json.Marshal(map[string]string{
		"message": "Transaction created successfully.",
	})
	return response, http.StatusCreated
}
