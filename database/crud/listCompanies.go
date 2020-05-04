package database

import (
	models "../models"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

//ListCompanies returns a list of all Companies.
//Returns an array of Companies.
func ListCompanies() []models.Company {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)

	if err != nil {
		log.Fatal(err)
		panic("failed to connect to database")
	}
	defer db.Close()

	var companies []models.Company
	db.Find(&companies)
	return companies
}
