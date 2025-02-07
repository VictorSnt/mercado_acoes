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

	transactions []models.Transaction,
	err error,
) {
	result := repo.Db.Where("user_id = ?", userId).Find(&transactions)

	if len(transactions) == 0 {
		return transactions, gorm.ErrRecordNotFound
	}

	return transactions, result.Error
}

func (repo TransactionsRepository) FindByEquitieId(equitieId uint) (

	transactions []models.Transaction,
	err error,
) {
	result := repo.Db.Where("equitie_id = ?", equitieId).Find(&transactions)
	return transactions, result.Error
}
