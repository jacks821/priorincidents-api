package database

import (
	"fmt"
	"os"

	models "github.com/jacks821/priorincidents-api/database/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//CreateLocation takes the arguments to make a Location and writes it to the database.
//Returns a Location and an error
func CreateLocation(streetNumber string, street string, city string, state string, zipCode string, storeNumber string, companyID string) (models.Location, error) {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
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
