package repositories

import (
	"mercado/acoes/database/models"
	DTO "mercado/acoes/dto"

	"gorm.io/gorm"
)

type UsersRepository struct {
	Db *gorm.DB
}

func (repo UsersRepository) Create(user DTO.CreateUser) error {
	result := repo.Db.Create(&models.User{
		Name:    user.Name,
		Balance: user.Balance,
	})

	return result.Error
}

func (repo UsersRepository) FindAll() (user []models.User, err error) {
	result := repo.Db.Find(&user)
	return user, result.Error
}

func (repo UsersRepository) FindById(id uint) (DTO.DisplayUser, error) {
	var user models.User
	result := repo.Db.First(&user, id)
	displayUser := DTO.DisplayUser{ID: user.ID, Name: user.Name, Balance: user.Balance}

	return displayUser, result.Error
}
