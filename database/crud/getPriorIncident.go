package database

import (
	models "../models"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

//GetPriorIncident finds a PriorIncident by its ID in the database.
//Returns a PriorIncident.
func GetPriorIncident(priorIncidentID string) models.PriorIncident {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)

	if err != nil {
		log.Fatal(err)
		panic("failed to connect to database")
	}
	defer db.Close()

	var priorIncident models.PriorIncident
	db.Preload("Reports").Where("id = ?", priorIncidentID).Find(&priorIncident)
	return priorIncident
}
