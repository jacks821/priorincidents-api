package database

import (
	"log"

	models "../models"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

//ListPriorIncidents returns a list of all PriorIncidents by LocationID.
//Returns an array of PriorIncidents.
func ListPriorIncidents(locationID string) []models.PriorIncident {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)

	if err != nil {
		log.Fatal(err)
		panic("failed to connect to database")
	}
	defer db.Close()

	var priorIncidents []models.PriorIncident
	db.Find(&priorIncidents, "location_id= ?", locationID)
	return priorIncidents
}
