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

// GetAllUser godoc
// @Summary Get all users
// @Description Get all users
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {array} DTO.DisplayUser
// @Failure 400 {object} map[string]string
// @Router /api/v1/users [get]
func GetAllUser(c *gin.Context) {
	var db *gorm.DB = database.GetConnection(configs.GetDbUri())
	response, status := handlers.FindAllUsers(db)
	c.JSON(status, response)
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} DTO.DisplayUser
// @Failure 400 {object} map[string]string
// @Router /api/v1/users/{id} [get]
func GetUser(c *gin.Context) {
	var db *gorm.DB = database.GetConnection(configs.GetDbUri())
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid UserID"})
		return
	}
	response, status := handlers.FindUserById(db, uint(idUint))
	c.JSON(status, response)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body DTO.CreateUser true "Create User"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/v1/users [post]
func CreateUser(c *gin.Context) {
	var db *gorm.DB = database.GetConnection(configs.GetDbUri())
	var newUser DTO.CreateUser
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "detail": "Invalid payload"})
		return
	}
	response, status := handlers.CreateUser(db, newUser)
	c.JSON(status, response)
}

// UpdateUser godoc
// @Summary Update a user by ID
// @Description Update a user by ID
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body DTO.UpdateUser true "Update User"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/v1/users/{id} [put]
func UpdateUser(c *gin.Context) {
	var db *gorm.DB = database.GetConnection(configs.GetDbUri())
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid UserID"})
		return
	}
	var updatedUser DTO.UpdateUser
	err = c.ShouldBindJSON(&updatedUser)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "detail": "Invalid payload"})
		return
	}
	response, status := handlers.UpdateUser(db, uint(idUint), updatedUser)
	c.JSON(status, response)
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/v1/users/{id} [delete]
func DeleteUser(c *gin.Context) {
	var db *gorm.DB = database.GetConnection(configs.GetDbUri())
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid UserID"})
		return
	}
	response, status := handlers.DeleteUser(db, uint(idUint))
	c.JSON(status, response)
}

// UserEquitieStock godoc
// @Summary Get user's equity stocks by user ID
// @Description Get user's equity stocks by user ID
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {array} DTO.DisplayUserEquities
// @Failure 400 {object} map[string]string
// @Router /api/v1/users/{id}/equitiestocks [get]
func UserEquitieStock(c *gin.Context) {
	var db *gorm.DB = database.GetConnection(configs.GetDbUri())
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid UserID"})
		return
	}
	response, status := handlers.RetriveUserEquitieStocks(db, uint(idUint))
	c.JSON(status, response)
}
