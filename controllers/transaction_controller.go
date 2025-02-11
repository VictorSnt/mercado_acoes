package controllers

import (
	"mercado/acoes/configs"
	"mercado/acoes/database"
	DTO "mercado/acoes/dto"
	"mercado/acoes/handlers"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllUserTransaction godoc
// @Summary Get all transactions for a user
// @Description Get all transactions for a user by user ID
// @Tags Transactions
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} []DTO.DisplayTransaction
// @Failure 400 {object} map[string]string
// @Router /api/v1/transactions/user/{id} [get]
func GetAllUserTransaction(c *gin.Context) {
	var db *gorm.DB = database.GetConnection(configs.GetDbUri())
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid UserID"})
		return
	}
	response, status := handlers.FindTransactionByUserId(db, uint(idUint))
	c.JSON(status, response)
}

// CreateTransaction godoc
// @Summary Create a new transaction
// @Description Create a new transaction with the input payload
// @Tags Transactions
// @Accept  json
// @Produce  json
// @Param transaction body DTO.CreateTransaction true "Transaction payload"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/v1/transactions [post]
func CreateTransaction(c *gin.Context) {
	var db *gorm.DB = database.GetConnection(configs.GetDbUri())
	var newTransaction DTO.CreateTransaction
	err := c.ShouldBindJSON(&newTransaction)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "detail": "Invalid payload"})
		return
	}
	response, status := handlers.CreateEquiteTransaction(db, newTransaction)
	c.JSON(status, response)
}
