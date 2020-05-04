package database

import (
	models "../models"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

//ListLocations returns a list of all Locations for a particular company by the Company ID.
//Returns an array of Locations.
func ListLocations(companyID string) ([]models.Location, error) {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)
	var locations []models.Location
	if err != nil {
		return locations, err
	}
	defer db.Close()

	db.Find(&locations, "company_id= ?", companyID)
	return locations, nil
}
