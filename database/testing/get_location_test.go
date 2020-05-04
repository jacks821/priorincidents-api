package database

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	crud "../crud"
)

func TestGetLocation(t *testing.T) {
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

	company := crud.CreateCompany("Meijer")

	location := crud.CreateLocation(streetNumber, street, city, state, zipCode, storeNumber, company.CommonModelFields.ID.String())

	result := crud.GetLocation(location.ID.String())

	if location.StreetNumber != result.StreetNumber {
		t.Error("expected", location.StreetNumber, "got", result.StreetNumber)
	}

	if location.Street != result.Street {
		t.Error("expected", location.Street, "got", result.Street)
	}

	if location.City != result.City {
		t.Error("expected", location.City, "got", result.City)
	}

	if location.State != result.State {
		t.Error("expected", location.State, "got", result.State)
	}

	if location.ZipCode != result.ZipCode {
		t.Error("expected", location.ZipCode, "got", result.ZipCode)
	}

	if location.StoreNumber != result.StoreNumber {
		t.Error("expected", location.StoreNumber, "got", result.StoreNumber)
	}

	if location.CompanyID != result.CompanyID {
		t.Error("expected", location.CompanyID, "got", result.CompanyID)
	}

	db.Delete(&company)
	db.Delete(&location)
}
