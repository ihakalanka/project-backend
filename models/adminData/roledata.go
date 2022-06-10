package adminData

import (
	"gorm.io/gorm"
)

type Role struct{
	gorm.Model
	Id int `gorm:"primaryKey" json:"id"`
	RoleName string `gorm:"unique" json:"rolename"`
}