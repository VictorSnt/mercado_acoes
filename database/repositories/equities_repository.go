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

func (repo EquitieRepository) FindAll() (equitiesDtoList []DTO.DisplayEquitie, err error) {
	var equities []models.Equitie
	result := repo.Db.Find(&equities)

	if len(equities) == 0 {
		return equitiesDtoList, gorm.ErrRecordNotFound
	}

	for _, equitie := range equities {
		equitieDTO := parseEquitieModelToDTO(equitie)
		equitiesDtoList = append(equitiesDtoList, equitieDTO)
	}

	return equitiesDtoList, result.Error
}

func (repo EquitieRepository) FindById(id uint) (DTO.DisplayEquitie, error) {
	var equitie models.Equitie
	result := repo.Db.First(&equitie, id)

	equitieDTO := parseEquitieModelToDTO(equitie)
	return equitieDTO, result.Error
}

func parseEquitieModelToDTO(equitie models.Equitie) DTO.DisplayEquitie {
	return DTO.DisplayEquitie{
		ID:                    equitie.ID,
		Name:                  equitie.Name,
		CurrentPrince:         equitie.CurrentPrince,
		PriceChangePercentage: equitie.CurrentPrince,
	}
}
