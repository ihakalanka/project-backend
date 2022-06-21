package customerData

import (
	"gorm.io/gorm"
)

type WishlistData struct{
	gorm.Model
	Id int `gorm:"primaryKey" json:"id"`
	ProductId int `json:"productid"`
	Imageurl        string `json:"imageurl"`
	Productprice    int    `json:"productprice"`
	ProductTitle    string `json:"producttitle"`
	ProductSubtitle string `json:"productsubtitle"`
	Quantity int    `json:"quantity"`
	UserId int `json:"userid"`
}