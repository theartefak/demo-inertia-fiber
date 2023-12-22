package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"     gorm:"type:text" validate:"required"`
	Email    string `json:"email"    gorm:"unique;"   validate:"required,email"`
	Password string `json:"password" gorm:"type:text" validate:"required,min=8"`
}

func (u User) MarshalJSON() ([]byte, error) {
    type Alias User
    return json.Marshal(&struct {
        Password string `json:"password"`
        *Alias
    }{
        Password : "********",
        Alias    : (*Alias)(&u),
    })
}
