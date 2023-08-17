package domain

import "gorm.io/gorm"

type Otp struct {
	gorm.Model
	Name   string `gorm:"column:name;not null" json:"name"`
	UserID uint   `gorm:"column:user_id;not null" json:"user_id"`
}

func (Otp) TableName() string {
	return "otp"
}
