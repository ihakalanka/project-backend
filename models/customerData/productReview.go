package customerData

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Id        int    `json:"id" gorm:"autoIncrement"`
	UserId    int    `json:"userId"`
	ProdId    int    `json:"prodId"`
	Name      string `json:"name"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
	UserEmail string `json:"userEmail"`
}
