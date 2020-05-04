package database

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	crud "github.com/jacks821/priorincidents-api/crud"
)

func TestCreateLocation(t *testing.T) {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)

	if err != nil {
		log.Fatal(err)
		panic("failed to connect to database")
	}
	defer db.Close()

	streetNumber := "4123"
	street := "Gibson Rd."
	city := "Henderson"
	state := "Nevada"
	zipCode := "89012"
	storeNumber := "1234"

	company, _ := crud.CreateCompany("Meijer")

	location, err := crud.CreateLocation(streetNumber, street, city, state, zipCode, storeNumber, company.CommonModelFields.ID.String())
	if err != nil {
		t.Error("Expected nil ", "got ", err)
	}
	if location.StreetNumber != streetNumber {
		t.Error("expected", streetNumber, "got", location.StreetNumber)
	}

	if location.Street != street {
		t.Error("expected", street, "got", location.Street)
	}

	if location.City != city {
		t.Error("expected", city, "got", location.City)
	}

	if location.State != state {
		t.Error("expected", state, "got", location.State)
	}

	if location.ZipCode != zipCode {
		t.Error("expected", zipCode, "got", location.ZipCode)
	}

	if location.StoreNumber != storeNumber {
		t.Error("expected", storeNumber, "got", location.StoreNumber)
	}

	if location.CompanyID != company.CommonModelFields.ID {
		t.Error("expected", company.CommonModelFields.ID, "got", location.CompanyID)
	}

	db.Delete(&company)
	db.Delete(&location)
}
