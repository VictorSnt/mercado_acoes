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

func (repo UsersRepository) FindAll() (userDtoList []DTO.DisplayUser, err error) {
	var users []models.User
	result := repo.Db.Find(&users)

	if len(users) == 0 {
		return userDtoList, gorm.ErrRecordNotFound
	}

	for _, user := range users {
		userDTO := parseUserModelToDTO(user)
		userDtoList = append(userDtoList, userDTO)
	}

	return userDtoList, result.Error
}

func (repo UsersRepository) FindById(id uint) (DTO.DisplayUser, error) {
	var user models.User
	result := repo.Db.First(&user, id)
	userDTO := parseUserModelToDTO(user)

	return userDTO, result.Error
}

func (repo UsersRepository) Update(id uint, user interface{}) error {
	statment := repo.Db.Model(&models.User{}).Where("id = ?", id)
	result := statment.Updates(user)
	return result.Error
}

func parseUserModelToDTO(user models.User) DTO.DisplayUser {
	return DTO.DisplayUser{
		ID:      user.ID,
		Name:    user.Name,
		Balance: user.Balance,
	}
}
