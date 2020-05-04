package database

import (
	"log"
	"time"

	models "../models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"os"
)

//CreatePriorIncident takes the arguments to make a PriorIncident and writes it to the database.
//Returns a PriorIncident
func CreatePriorIncident(date string, fallType string, attorneyName string, locationID string) models.PriorIncident {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)

	if err != nil {
		log.Fatal(err)
		panic("failed to connect to database")
	}
	defer db.Close()

	dateObject, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Fatal(err)
	}

	locID, err := uuid.FromString(locationID)
	if err != nil {
		log.Fatal(err)
	}

	priorIncident := models.PriorIncident{Date: dateObject, FallType: fallType, AttorneyName: attorneyName, LocationID: locID}
	db.NewRecord(priorIncident)
	db.Create(&priorIncident)
	return priorIncident
}
