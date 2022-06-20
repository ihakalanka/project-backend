package privilegeData

import (
	"gorm.io/gorm"
)

type Privilegedata struct {
	gorm.Model
	Id                        int    `gorm:"primaryKey" json:"id"`
	PrivilegeName        	  string `json:"privilegename"`
	PrivilegeDescription      string `json:"privilegedescription"`                    
}