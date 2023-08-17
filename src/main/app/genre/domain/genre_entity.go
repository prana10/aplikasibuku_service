package domain

import (
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	Name string `gorm:"unique;not null;column:name" json:"name"`
}

func (Genre) TableName() string {
	return "genres"
}
