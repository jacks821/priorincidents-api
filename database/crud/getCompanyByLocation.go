package database

import (
	models "../models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"os"
)

//GetCompanyByLocation finds a Company by the ID of one of its Locations in the database.
//Returns a Company.
func GetCompanyByLocation(locationID string) (models.Company, error) {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)
	var company models.Company
	if err != nil {
		return company, err
	}
	defer db.Close()

	locID, err := uuid.FromString(locationID)
	if err != nil {
		return company, err
	}

	var location models.Location
	db.Where("id = ?", locID).Find(&location)
	db.Where("id = ?", location.CompanyID).Find(&company)
	return company, nil
}
