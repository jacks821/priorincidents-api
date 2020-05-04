package database

import (
	models "../models"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

//GetCompany finds a Company by its ID in the database.
//Returns a Company.
func GetCompany(companyID string) models.Company {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)

	if err != nil {
		log.Fatal(err)
		panic("failed to connect to database")
	}
	defer db.Close()

	var company models.Company
	db.Preload("Locations").Where("id = ?", companyID).Find(&company)
	return company
}
