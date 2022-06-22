package companyControllers

import (

	"github.com/gofiber/fiber/v2"

	"main.go/database"
	"main.go/models/companyData"
)

func Getcat(c *fiber.Ctx) error {
	db := database.DB
	var company []CompanyData.CompanyData
	db.Find(&company)
	
	return c.JSON(company)
}

func Postcat(c *fiber.Ctx) error {
		db := database.DB
		var company  CompanyData.CompanyData

		err := c.BodyParser(&company); 
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Error",
				"data":    err,
			})
		}
		var company1 CompanyData.CompanyData
		name := company.CompanyName
		err = db.Find(&company1, "company_name = ?", name).Error
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "error in delete company",
			})
		}
		
		if name == company1.CompanyName {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Duplicate role available for company",
			})
		}
		db.Create(&company)
		return c.JSON(company)
	}

func Getcatid(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var company CompanyData.CompanyData
	db.Find(&company, id)
	return c.JSON(company)
}	

func Deletecat(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var company CompanyData.CompanyData
	err := db.Find(&company, "id = ?", id).Error
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "error in delete company",
			})
		}
	db.Unscoped().Delete(&company)
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Company data deleted",
		})	
}
