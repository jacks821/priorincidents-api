package database

import (
	models "../models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"log"
	"os"
)

//GetCompanyByLocation finds a Company by the ID of one of its Locations in the database.
//Returns a Company.
func GetCompanyByLocation(locationID string) models.Company {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)

	if err != nil {
		log.Fatal(err)
		panic("failed to connect to database")
	}
	defer db.Close()

	locID, err := uuid.FromString(locationID)
	if err != nil {
		log.Fatal(err)
	}

	var location models.Location
	db.Where("id = ?", locID).Find(&location)
	var company models.Company
	db.Where("id = ?", location.CompanyID).Find(&company)
	return company
}
