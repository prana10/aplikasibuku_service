package users

import (
	domainUser "service-api/src/main/app/users/domain"
)

type UserFormatter struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Token       string `json:"token"`
}

func FormatUser(user domainUser.User, token string) UserFormatter {
	var userFormatter UserFormatter

	userFormatter.ID = user.ID
	userFormatter.Email = user.Email
	userFormatter.Username = user.Username
	userFormatter.PhoneNumber = user.PhoneNumber
	userFormatter.Address = user.Address
	userFormatter.Token = token

	return userFormatter
}

func FormatUsers(users []domainUser.User) []UserFormatter {
	var usersFormatter []UserFormatter

	for _, user := range users {
		userFormatter := FormatUser(user, "")
		usersFormatter = append(usersFormatter, userFormatter)
	}

	return usersFormatter
}
