package impl

import (
	domainUser "service-api/src/main/app/users/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) FindAll() ([]domainUser.User, error) {
	var users []domainUser.User
	error := repo.db.Find(&users).Error
	if error != nil {
		return users, error
	}

	return users, nil
}
func (repo *userRepository) CreateUser(user domainUser.User) (domainUser.User, error) {
	error := repo.db.Create(&user).Error
	if error != nil {
		return user, error
	}

	return user, nil
}
func (repo *userRepository) FindUserByID(id uint) (domainUser.User, error) {
	var user domainUser.User
	error := repo.db.First(&user, id).Error
	if error != nil {
		return user, error
	}

	return user, nil
}

func (repo *userRepository) FindUserByEmail(email string) (domainUser.User, error) {
	var user domainUser.User
	error := repo.db.Where("email = ?", email).First(&user).Error
	if error != nil {
		return user, error
	}

	return user, nil
}
