package database

import (
	"os"

	models "github.com/jacks821/priorincidents-api/database/models"
	"github.com/jinzhu/gorm"
)

//CreateCompany takes the arguments to make a Company and writes it to the database.
//Returns a Company and an error
func CreateCompany(name string) (models.Company, error) {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	company := models.Company{Name: name, Locations: []models.Location{}}
	if err != nil {
		return company, err
	}
	defer db.Close()

	db.NewRecord(company)
	db.Create(&company)
	return company, nil
}
