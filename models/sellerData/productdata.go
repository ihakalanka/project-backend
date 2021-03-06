package sellerData

import (
	"gorm.io/gorm"
)

type Productdata struct {
	gorm.Model
	Id              int    `gorm:"primaryKey" json:"id"`
	ProductTitle    string `json:"producttitle"`
	ProductSubtitle string `json:"productsubtitle"`
	CategoryName    string `json:"categoryname"`
	Imageurl        string `json:"imageurl"`
	Description     string `json:"description"`
	Productprice    int    `json:"productprice"`
	Productquantity int    `json:"productquantity"`
	UserId          int    `json:"userid"`
	AverageRate     int    `json:"averagerate"`
}
