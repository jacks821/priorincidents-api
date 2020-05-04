package database

import (
	"fmt"
	"os"

	models "github.com/jacks821/priorincidents-api/database/models"
	"github.com/jinzhu/gorm"
)

//ListPriorIncidents returns a list of all PriorIncidents by LocationID.
//Returns an array of PriorIncidents and an error.
func ListPriorIncidents(locationID string) ([]models.PriorIncident, error) {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)
	var priorIncidents []models.PriorIncident
	if err != nil {
		return priorIncidents, err
	}
	defer db.Close()

	db.Find(&priorIncidents, "location_id= ?", locationID)
	return priorIncidents, nil
}
