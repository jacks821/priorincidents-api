package database

import (
	"fmt"
	"os"

	models "github.com/jacks821/priorincidents-api/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//CreateLocation takes the arguments to make a Location and writes it to the database.
//Returns a Location and an error
func CreateLocation(streetNumber string, street string, city string, state string, zipCode string, storeNumber string, companyID string) (models.Location, error) {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)
	location := models.Location{StreetNumber: streetNumber, Street: street, City: city, State: state, ZipCode: zipCode, StoreNumber: storeNumber}
	if err != nil {
		return location, err
	}
	defer db.Close()

	coID, err := uuid.FromString(companyID)
	if err != nil {
		return location, err
	}
	location.CompanyID = coID

	db.NewRecord(location)
	db.Create(&location)
	return location, nil
}
