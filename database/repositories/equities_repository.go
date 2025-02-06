package repositories

import (
	"mercado/acoes/database/models"
	DTO "mercado/acoes/dto"

	"gorm.io/gorm"
)

type EquitieRepository struct {
	Db *gorm.DB
}

func (repo EquitieRepository) Create(equitie DTO.CreateEquitie) error {
	result := repo.Db.Create(
		&models.Equitie{
			Name:                  equitie.Name,
			CurrentPrince:         equitie.CurrentPrince,
			PriceChangePercentage: equitie.PriceChangePercentage,
		},
	)

	return result.Error
}

func (repo EquitieRepository) FindAll() (equities []models.Equitie, err error) {
	result := repo.Db.Find(&equities)
	return equities, result.Error
}

func (repo EquitieRepository) FindById(id uint) (equitie models.Equitie, err error) {
	result := repo.Db.First(&equitie, id)
	return equitie, result.Error
}
