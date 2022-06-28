package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id        uint   `json:"id" gorm:"autoIncrement"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `gorm:"unique" json:"email" validate:"email,required"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

type PasswordReset struct {
	gorm.Model
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Token string `gorm:"unique"`
}
