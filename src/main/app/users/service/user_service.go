package service

import (
	domainUser "service-api/src/main/app/users/domain"
	inputUser "service-api/src/main/util/input/users"
)

type UserService interface {
	FindAllUser() ([]domainUser.User, error)
	LoginUser(user inputUser.LoginUserInput) (domainUser.User, error)
	RegisterUser(user inputUser.RegisterUserInput) (domainUser.User, error)
	GetUserByID(id uint) (domainUser.User, error)
}
