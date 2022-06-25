package merchantApplicationData

import (
	"gorm.io/gorm"
)

type MerchantApplicationdata struct {
	gorm.Model
	Id                        int    `gorm:"primaryKey" json:"id"`
	MerchantLegalName         string `json:"merchantlegalname"`
	MerchantNameOnCardStatement string `json:"merchantnameoncardstatement"`
	OfficialWebsite           string `json:"officialwebsite"`
	ContactPersonName         string `json:"contactpersonname"`
	ContactPersonEmailID      string `json:"contactpersonemailid"`
	ContactPersonMobileNumber string `json:"contactpersonmobilenumber"`
	BusinessAddress           string `json:"businessaddress"`
	Province                  string `json:"province"`
	District                  string `json:"district"`
	PostelCode                string `json:"postelcode"`
	Email                     string `json:"email"`
	Tel                       string `json:"tel"`
	Fax                       string `json:"fax"`
	ProductDescription        string `json:"productdescription"`
	YearsOfIncorporation      int    `json:"yearsofincorporation"`
}
