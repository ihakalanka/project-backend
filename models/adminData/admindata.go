package adminData

import (
	"gorm.io/gorm"
)

type AdminDetails struct{
	gorm.Model
	Id int `json:"id"`
	Name string `json:"adminname"`
	Email string `json:"adminemail"`
	Telephone string `json:"telephone"`
	Gender string `json:"gender"`
	ProfileImageurl string `json:"profileimgeurl"`
}