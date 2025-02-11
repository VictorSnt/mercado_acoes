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

// GetAllEquitie godoc
// @Summary Get all equities
// @Description Get all equities
// @Tags Equities
// @Accept  json
// @Produce  json
// @Success 200 {array} DTO.DisplayEquitie
// @Failure 500 {object} map[string]string
// @Router /api/v1/equities [get]
func GetAllEquitie(c *gin.Context) {
	var db *gorm.DB = database.GetConnection(configs.GetDbUri())
	response, status := handlers.FindAllEquities(db)
	c.JSON(status, response)
}

// GetEquitie godoc
// @Summary Get an equitie by ID
// @Description Get an equitie by ID
// @Tags Equities
// @Accept  json
// @Produce  json
// @Param id path int true "Equitie ID"
// @Success 200 {object} DTO.DisplayEquitie
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/equities/{id} [get]
func GetEquitie(c *gin.Context) {
	var db *gorm.DB = database.GetConnection(configs.GetDbUri())
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid EquitieID"})
		return
	}
	response, status := handlers.FindEquitieById(db, uint(idUint))
	c.JSON(status, response)
}

// CreateEquitie godoc
// @Summary Create a new equitie
// @Description Create a new equitie
// @Tags Equities
// @Accept  json
// @Produce  json
// @Param equitie body DTO.CreateEquitie true "Create Equitie"
// @Success 201 {object} DTO.DisplayEquitie
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/equities [post]
func CreateEquitie(c *gin.Context) {
	var db *gorm.DB = database.GetConnection(configs.GetDbUri())
	var newEquitie DTO.CreateEquitie
	err := c.ShouldBindJSON(&newEquitie)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "detail": "Invalid payload"})
		return
	}
	response, status := handlers.CreateEquitie(db, newEquitie)
	c.JSON(status, response)
}

// UpdateEquitie godoc
// @Summary Update an existing equitie
// @Description Update an existing equitie
// @Tags Equities
// @Accept  json
// @Produce  json
// @Param id path int true "Equitie ID"
// @Param equitie body DTO.UpdateEquitie true "Update Equitie"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/equities/{id} [put]
func UpdateEquitie(c *gin.Context) {
	var db *gorm.DB = database.GetConnection(configs.GetDbUri())
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	var updatedEquitie DTO.UpdateEquitie
	err = c.ShouldBindJSON(&updatedEquitie)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "detail": "Invalid payload"})
		return
	}
	response, status := handlers.UpdateEquitieName(db, uint(idUint), updatedEquitie)
	c.JSON(status, response)
}
