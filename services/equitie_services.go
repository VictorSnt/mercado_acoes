package services

import (
	"errors"
	"fmt"
	"mercado/acoes/database/repositories"
	DTO "mercado/acoes/dto"
	"mercado/acoes/enums"

	"gorm.io/gorm"
)

func UpdateEquitiePrice(
	db *gorm.DB,
	equitieId uint,
	transactionType enums.TransactionType,

) (DTO.UpdateEquitiePrice, error) {

	EquitiesRepository := repositories.EquitiesRepository{Db: db}
	equitie, err := EquitiesRepository.FindById(equitieId)
	if err != nil {
		return DTO.UpdateEquitiePrice{}, err
	}

	if transactionType == enums.TransactionBuyOperation {
		currentPrice := equitie.CurrentPrice + (equitie.CurrentPrice * equitie.PriceChangePercentage / 100)
		return DTO.UpdateEquitiePrice{CurrentPrice: currentPrice}, nil
	}

	if transactionType == enums.TransactionSaleOperation {
		currentPrice := equitie.CurrentPrice - (equitie.CurrentPrice * equitie.PriceChangePercentage / 100)
		return DTO.UpdateEquitiePrice{CurrentPrice: currentPrice}, nil
	}

	return DTO.UpdateEquitiePrice{}, errors.New("invalid transaction type")
}

func ValidadeUserEquitieTransaction(
	db *gorm.DB,
	newTransaction DTO.CreateTransaction,
) error {
	UsersRepository := repositories.UsersRepository{Db: db}
	EquitiesRepository := repositories.EquitiesRepository{Db: db}
	TransactionsRepository := repositories.TransactionsRepository{Db: db}

	user, err := UsersRepository.FindById(newTransaction.UserID)

	if err != nil {
		return fmt.Errorf("user not found: %v", err)
	}

	equitie, err := EquitiesRepository.FindById(newTransaction.EquitieID)

	if err != nil {
		return fmt.Errorf("equitie not found: %v", err)
	}

	if newTransaction.Type == string(enums.TransactionBuyOperation) {
		if user.Balance < (equitie.CurrentPrice * float64(newTransaction.Quantity)) {
			return errors.New("insufficient balance")
		}
		return nil
	}

	if newTransaction.Type == string(enums.TransactionSaleOperation) {
		var userEquitiesCount uint

		transactionsList, err := TransactionsRepository.FindByUserIdAndEquitieId(
			newTransaction.UserID,
			newTransaction.EquitieID,
		)

		if err != nil {
			return fmt.Errorf("error while search user transactions: %v", err)
		}

		for _, transaction := range transactionsList {
			if transaction.Type == string(enums.TransactionBuyOperation) {
				userEquitiesCount += transaction.Quantity
			}

			if transaction.Type == string(enums.TransactionSaleOperation) {
				userEquitiesCount -= transaction.Quantity
			}
		}

		if userEquitiesCount < newTransaction.Quantity {
			return errors.New("insufficient equities for sale operation")
		}

		return nil
	}

	return errors.New("invalid transaction type in the new transaction")
}

func UpdateUserBalance(
	db *gorm.DB,
	userId uint,
	transactionCost float64,
	transactionType string,
) error {
	UsersRepository := repositories.UsersRepository{Db: db}
	user, err := UsersRepository.FindById(userId)

	if err != nil {
		return fmt.Errorf("user not found: %v", err)
	}

	switch transactionType {

	case string(enums.TransactionBuyOperation):
		if user.Balance < transactionCost {
			return errors.New("insufficient balance")
		}
		user.Balance -= transactionCost

	case string(enums.TransactionSaleOperation):
		user.Balance += transactionCost

	default:
		return errors.New("invalid transaction type")
	}

	err = UsersRepository.Update(user.ID, DTO.UpdateUserBalance{Balance: user.Balance})

	if err != nil {
		return fmt.Errorf("error while update user balance: %v", err)
	}

	return nil
}
