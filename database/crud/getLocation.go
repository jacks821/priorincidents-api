package database

import (
	"os"

	models "github.com/jacks821/priorincidents-api/database/models"
	"github.com/jinzhu/gorm"
)

//GetLocation finds a Location by its ID in the database.
//Returns a Location.
func GetLocation(locationID string) (models.Location, error) {
	//s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	var location models.Location
	if err != nil {
		return location, err
	}
	defer db.Close()

	db.Preload("PriorIncidents").Preload("Reports").Where("id = ?", locationID).Find(&location)
	return location, nil
}
