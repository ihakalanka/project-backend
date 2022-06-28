package AddressData

import (
	"gorm.io/gorm"	 
)

type AddressData struct {
	gorm.Model
	ShippingID       int    `gorm:"primaryKey" json:"id"`
	ReceiverName     string `json:"receivername"`
	StreetLine01    string `json:"streetline01"`
	StreetLine02    string `json:"streetline02"`
	City    string `json:"city"`
	Country  string `json:"country"`
	ZipCode string `json:"zipcode"`
}
