package database

import (
	"log"

	models "../models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"os"
)

//CreateLocation takes the arguments to make a Location and writes it to the database.
//Returns a Location
func CreateLocation(streetNumber string, street string, city string, state string, zipCode string, storeNumber string, companyID string) models.Location {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)
	if err != nil {
		log.Fatal(err)
		panic("failed to connect to database")
	}
	defer db.Close()

	coID, err := uuid.FromString(companyID)
	if err != nil {
		log.Fatal(err)
	}

	location := models.Location{StreetNumber: streetNumber, Street: street, City: city, State: state, ZipCode: zipCode, StoreNumber: storeNumber, CompanyID: coID}
	db.NewRecord(location)
	db.Create(&location)
	return location
}
