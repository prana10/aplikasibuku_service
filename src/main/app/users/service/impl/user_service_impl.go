package impl

import (
	"errors"
	domainUser "service-api/src/main/app/users/domain"
	repositoryUser "service-api/src/main/app/users/repository"
	inputUser "service-api/src/main/util/input/users"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repositoryUser.UserRepository
}

func NewUserService(repoUser repositoryUser.UserRepository) *userService {
	return &userService{userRepo: repoUser}
}

func (service *userService) FindAllUser() ([]domainUser.User, error) {
	users, err := service.userRepo.FindAll()

	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return users, errors.New("No user found")
	}

	return users, nil
}

func (service *userService) GetUserByID(id uint) (domainUser.User, error) {
	user, err := service.userRepo.FindUserByID(id)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	return user, nil
}

func (service *userService) LoginUser(userInput inputUser.LoginUserInput) (domainUser.User, error) {
	email := userInput.Email
	password := userInput.Password

	user, err := service.userRepo.FindUserByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (service *userService) RegisterUser(userInput inputUser.RegisterUserInput) (domainUser.User, error) {
	user := domainUser.User{
		Name:        userInput.Name,
		Username:    userInput.Username,
		Email:       userInput.Email,
		PhoneNumber: userInput.PhoneNumber,
		Address:     userInput.Address,
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	user, errRepo := service.userRepo.CreateUser(user)
	if errRepo != nil {
		return user, errRepo
	}

	return user, nil
}
