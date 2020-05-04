package database

import (
	"time"

	"fmt"
	"os"

	models "github.com/jacks821/priorincidents-api/database/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//CreatePriorIncident takes the arguments to make a PriorIncident and writes it to the database.
//Returns a PriorIncident and an error
func CreatePriorIncident(date string, fallType string, attorneyName string, locationID string) (models.PriorIncident, error) {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)

	priorIncident := models.PriorIncident{FallType: fallType, AttorneyName: attorneyName}

	if err != nil {
		return priorIncident, err
	}
	defer db.Close()

	dateObject, err := time.Parse("2006-01-02", date)
	if err != nil {
		return priorIncident, err
	}
	priorIncident.Date = dateObject

	locID, err := uuid.FromString(locationID)
	if err != nil {
		return priorIncident, err
	}
	priorIncident.LocationID = locID

	db.NewRecord(priorIncident)
	db.Create(&priorIncident)
	return priorIncident, nil
}
