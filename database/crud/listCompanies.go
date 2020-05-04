package database

import (
	"fmt"
	"os"

	models "github.com/jacks821/priorincidents-api/models"
	"github.com/jinzhu/gorm"
)

//ListCompanies returns a list of all Companies.
//Returns an array of Companies and an error if applicable.
func ListCompanies() ([]models.Company, error) {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)
	var companies []models.Company
	if err != nil {
		return companies, err
	}
	defer db.Close()

	db.Find(&companies)
	return companies, nil
}
