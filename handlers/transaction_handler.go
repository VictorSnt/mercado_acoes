package handlers

import (
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
) (map[string]string, int) {

	TransactionRepository := repositories.TransactionsRepository{Db: db}
	EquitiesRepository := repositories.EquitiesRepository{Db: db}

	err := services.ValidadeUserEquitieTransaction(db, newTransaction)
	if err != nil {
		errorResponse := map[string]string{
			"error": err.Error(),
			"detail": fmt.Sprintf(
				"invalid transaction for userID %d.",
				newTransaction.UserID,
			),
		}
		return errorResponse, http.StatusBadRequest
	}

	equitieDTO, err := EquitiesRepository.FindById(newTransaction.EquitieID)
	if err != nil {
		errorResponse := map[string]string{
			"error": err.Error(),
			"detail": fmt.Sprintf(
				"error while finding equitie %d.",
				newTransaction.EquitieID,
			),
		}
		return errorResponse, http.StatusBadRequest
	}

	var transactionValue float64 = float64(newTransaction.Quantity) * equitieDTO.CurrentPrice
	err = services.UpdateUserBalance(
		db, newTransaction.UserID,
		transactionValue,
		newTransaction.Type,
	)
	if err != nil {
		errorResponse := map[string]string{
			"error": err.Error(),
			"detail": fmt.Sprintf(
				"error while charging user %d.",
				newTransaction.UserID,
			),
		}
		return errorResponse, http.StatusBadRequest
	}

	err = TransactionRepository.Create(newTransaction)
	if err != nil {
		errorResponse := map[string]string{"error": err.Error()}
		return errorResponse, http.StatusBadRequest
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
		return errorResponse, http.StatusBadRequest
	}

	err = EquitiesRepository.Update(newTransaction.EquitieID, equiteUpdateDTO)
	if err != nil {
		errorResponse := map[string]string{
			"error":  err.Error(),
			"detail": "Error while updating equitie price.",
		}
		return errorResponse, http.StatusBadRequest
	}

	successResponse := map[string]string{
		"message": "Transaction created successfully.",
	}
	return successResponse, http.StatusCreated
}

func FindTransactionByUserId(db *gorm.DB, id uint) (interface{}, int) {
	TransactionRepository := repositories.TransactionsRepository{Db: db}
	transaction, err := TransactionRepository.FindByUserId(id)
	if err != nil {
		errorResponse := map[string]string{
			"error":   err.Error(),
			"details": fmt.Sprintf("erro finding transaction by userID %d", id),
		}
		return errorResponse, http.StatusNotFound
	}

	return transaction, http.StatusOK
}
