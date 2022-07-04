package CompanyData

import (
	"gorm.io/gorm"
)

type CompanyData struct {
	gorm.Model
	CompanyID      int    `gorm:"primaryKey" json:"id"`
	CompanyName    string `json:"companyname"`
	EmailAddress   string `json:"emailaddress"`
	CreatePassword string `json:"createpassword"`
}
