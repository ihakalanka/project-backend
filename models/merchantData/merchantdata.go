package merchantData

import (
	"gorm.io/gorm"
)

type Merchantdata struct {
	gorm.Model
	Id                        int    `gorm:"primaryKey" json:"id"`
	MerchantLegalName         string `json:"merchantlegalname"`
	OfficialWebsite           string `json:"officialwebsite"`
	ContactPersonEmailID      string `json:"contactpersonemailid"`
	ContactPersonMobileNumber string `json:"contactpersonmobilenumber"`
	BusinessAddress           string `json:"businessaddress"`
	Profile                   string `json:"profile"`
	ProductDescription        string `json:"productdescription"`
	AverageProductValue       int    `json:"averageproductvalue"`
	CompanyLogourl            string `json:"companylogourl"`
}
