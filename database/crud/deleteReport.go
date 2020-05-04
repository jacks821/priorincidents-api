package database

import (
	"fmt"
	"os"

	models "github.com/jacks821/priorincidents-api/database/models"
	"github.com/jinzhu/gorm"
)

//DeleteReport takes the arguments to delete a Report and soft deletes it from the database.
//Returns an error.
func DeleteReport(id string) error {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)

	if err != nil {
		return err
	}
	defer db.Close()

	db.Where("id = ?", id).Delete(&models.Report{})
	return nil
}
