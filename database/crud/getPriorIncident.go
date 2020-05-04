package database

import (
	models "../models"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

//GetPriorIncident finds a PriorIncident by its ID in the database.
//Returns a PriorIncident.
func GetPriorIncident(priorIncidentID string) (models.PriorIncident, error) {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)
	var priorIncident models.PriorIncident
	if err != nil {
		return priorIncident, err
	}
	defer db.Close()

	db.Preload("Reports").Where("id = ?", priorIncidentID).Find(&priorIncident)
	return priorIncident, nil
}
