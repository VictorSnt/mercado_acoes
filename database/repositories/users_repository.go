package repositories

import (
	"mercado/acoes/database/models"
	DTO "mercado/acoes/dto"

	"gorm.io/gorm"
)

type UsuarioRepository struct {
	Db *gorm.DB
}

func (repo UsuarioRepository) Create(user DTO.CreateUser) error {
	result := repo.Db.Create(&models.Usuario{
		Name:    user.Name,
		Balance: user.Balance,
	})

	return result.Error
}

func (repo UsuarioRepository) FindAll() (user []models.Usuario, err error) {
	result := repo.Db.Find(&user)
	return user, result.Error
}

func (repo UsuarioRepository) FindById(id uint) (DTO.DisplayUser, error) {
	var user models.Usuario
	result := repo.Db.First(&user, id)
	displayUser := DTO.DisplayUser{ID: user.ID, Name: user.Name, Balance: user.Balance}

	return displayUser, result.Error
}
