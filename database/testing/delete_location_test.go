package database

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	crud "github.com/jacks821/priorincidents-api/database/crud"
)

func TestDeleteLocation(t *testing.T) {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)

	if err != nil {
		log.Fatal(err)
		panic("failed to connect to database")
	}
	defer db.Close()

	name := "Meijer"
	company, err := crud.CreateCompany(name)
	if err != nil {
		t.Error("Expected nil ", "got ", err)
	}
	streetNumber := "4123"
	street := "Gibson Rd."
	city := "Henderson"
	state := "Nevada"
	zipCode := "89012"
	storeNumber := "1234"

	location, _ := crud.CreateLocation(streetNumber, street, city, state, zipCode, storeNumber, company.CommonModelFields.ID.String())

	err = crud.DeleteLocation(location.ID.String())
	if err != nil {
		t.Error("Expected nil ", "got ", err)
	}
	err = crud.DeleteCompany(company.ID.String())
	if err != nil {
		t.Error("Expected nil ", "got ", err)
	}
}
