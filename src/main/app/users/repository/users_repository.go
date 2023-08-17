package repository

import (
	domainUser "service-api/src/main/app/users/domain"
)

type UserRepository interface {
	FindAll() ([]domainUser.User, error)
	FindUserByID(id uint) (domainUser.User, error)
	CreateUser(user domainUser.User) (domainUser.User, error)
	FindUserByEmail(email string) (domainUser.User, error)
}
