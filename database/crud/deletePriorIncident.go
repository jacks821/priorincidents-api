package database

import (
	models "../models"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

//DeletePriorIncident takes the arguments to delete a Prior Incident and soft deletes it from the database.
//Returns an error.
func DeletePriorIncident(id string) error {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)

	if err != nil {
		return err
	}
	defer db.Close()

	db.Where("id = ?", id).Delete(&models.PriorIncident{})
	return nil
}
