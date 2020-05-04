package database

import (
	"fmt"
	"os"

	models "github.com/jacks821/priorincidents-api/models"
	"github.com/jinzhu/gorm"
)

//CreateCompany takes the arguments to make a Company and writes it to the database.
//Returns a Company and an error
func CreateCompany(name string) (models.Company, error) {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)
	company := models.Company{Name: name, Locations: []models.Location{}}
	if err != nil {
		return company, err
	}
	defer db.Close()

	db.NewRecord(company)
	db.Create(&company)
	return company, nil
}
