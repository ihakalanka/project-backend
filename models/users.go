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
	Previlage string `json:"previlage"`
}

type PasswordReset struct {
	Id    uint
	Email string
	Token string `gorm:"unique"`
}

type Count struct{
	CountUser int `json:"countuser"`
}