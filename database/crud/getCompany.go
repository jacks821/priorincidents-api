package database

import (
	models "../models"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

//GetCompany finds a Company by its ID in the database.
//Returns a Company and an error
func GetCompany(companyID string) (models.Company, error) {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)
	var company models.Company
	if err != nil {
		return company, err
	}
	defer db.Close()
	db.Preload("Locations").Where("id = ?", companyID).Find(&company)
	return company, nil
}
