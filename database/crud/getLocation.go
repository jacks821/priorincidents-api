package database

import (
	models "../models"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

//GetLocation finds a Location by its ID in the database.
//Returns a Location.
func GetLocation(locationID string) models.Location {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)

	if err != nil {
		log.Fatal(err)
		panic("failed to connect to database")
	}
	defer db.Close()

	var location models.Location
	db.Preload("PriorIncidents").Preload("Reports").Where("id = ?", locationID).Find(&location)
	return location
}
