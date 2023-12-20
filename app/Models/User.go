package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name" gorm:"type:text"`
	Email string `json:"email" gorm:"unique;"`
	Password string `json:"-" gorm:"type:text"`
}
