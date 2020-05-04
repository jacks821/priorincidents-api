package database

import (
	models "../models"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

//CreateCompany takes the arguments to make a Company and writes it to the database.
//Returns a Company
func CreateCompany(name string) models.Company {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)

	if err != nil {
		log.Fatal(err)
		panic("failed to connect to database")
	}
	defer db.Close()

	company := models.Company{Name: name, Locations: []models.Location{}}
	db.NewRecord(company)
	db.Create(&company)
	return company
}
