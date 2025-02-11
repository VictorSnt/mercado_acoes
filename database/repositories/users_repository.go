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
	result := repo.Db.Where("deleted_at IS NULL").Find(&users)

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

func (repo UsersRepository) Delete(id uint) error {
	result := repo.Db.Delete(&models.User{}, "id = ?", id)
	return result.Error
}

func (repo UsersRepository) FindEquities(userID uint) (
	userStocks []DTO.DisplayUserEquities,
	err error,
) {
	subQuery := repo.Db.Model(&models.Transaction{}).
		Select(`
        transactions.user_id, 
        transactions.equitie_id, 
        equities.current_price, 
        SUM(CASE WHEN transactions.type = ? THEN transactions.quantity ELSE -transactions.quantity END) AS total_quantity`,
			"BUY").
		Joins(`
        JOIN equities ON equities.id = transactions.equitie_id`).
		Where("transactions.user_id = ?", userID).
		Group(`
        transactions.equitie_id, 
        equities.current_price`).
		Having(`
        SUM(CASE WHEN transactions.type = ? THEN transactions.quantity ELSE -transactions.quantity END) > 0`,
			"BUY")

	err = repo.Db.Table("(?) as stock_totals", subQuery).
		Select(`
		user_id,
        equitie_id, 
        current_price, 
        total_quantity, 
        current_price * total_quantity AS equitie_total_value`).
		Find(&userStocks).Error

	return userStocks, err
}
