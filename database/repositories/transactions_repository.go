package repositories

import (
	"mercado/acoes/database/models"
	DTO "mercado/acoes/dto"

	"gorm.io/gorm"
)

type TransactionsRepository struct {
	Db *gorm.DB
}

func (repo TransactionsRepository) Create(transactions DTO.CreateTransaction) error {
	result := repo.Db.Create(
		&models.Transaction{
			UserID:          transactions.UserID,
			EquitieID:       transactions.EquitieID,
			Type:            transactions.Type,
			Quantity:        transactions.Quantity,
			UnitPrice:       transactions.UnitPrice,
			TransactionDate: transactions.TransactionDate,
		},
	)

	return result.Error
}

func (repo TransactionsRepository) FindByUserId(userId uint) (

	transactionsDtoList []DTO.DisplayTransaction,
	err error,
) {
	var transactions []models.Transaction
	statement := repo.Db.Where("user_id = ?", userId)

	result := statement.Find(&transactions)
	if len(transactions) == 0 {
		return transactionsDtoList, gorm.ErrRecordNotFound
	}

	for _, transaction := range transactions {
		transactionsDTO := parseTransactionModelToDTO(transaction)
		transactionsDtoList = append(transactionsDtoList, transactionsDTO)
	}

	return transactionsDtoList, result.Error
}

func (repo TransactionsRepository) FindByEquitieId(equitieId uint) (

	transactionsDtoList []DTO.DisplayTransaction,
	err error,
) {
	var transactions []models.Transaction
	result := repo.Db.Where("equitie_id = ?", equitieId).Find(&transactions)

	for _, transaction := range transactions {
		transactionsDTO := parseTransactionModelToDTO(transaction)
		transactionsDtoList = append(transactionsDtoList, transactionsDTO)
	}

	return transactionsDtoList, result.Error
}

func parseTransactionModelToDTO(transaction models.Transaction) DTO.DisplayTransaction {
	return DTO.DisplayTransaction{
		ID:              transaction.ID,
		UserID:          transaction.UserID,
		EquitieID:       transaction.EquitieID,
		Type:            transaction.Type,
		Quantity:        transaction.Quantity,
		UnitPrice:       transaction.UnitPrice,
		TransactionDate: transaction.TransactionDate,
	}
}
