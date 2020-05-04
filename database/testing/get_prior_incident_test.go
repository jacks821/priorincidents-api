package database

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	crud "../crud"
)

func TestGetPriorIncident(t *testing.T) {
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

	location, _ := crud.CreateLocation(streetNumber, street, city, state, zipCode, storeNumber, company.CommonModelFields.ID.String())

	date := time.Now()
	createdDate := date.Format("2006-01-02")
	fallType := "Slip"
	attorneyName := "Charlie Jackson"

	priorIncident, _ := crud.CreatePriorIncident(createdDate, fallType, attorneyName, location.CommonModelFields.ID.String())

	result, err := crud.GetPriorIncident(priorIncident.ID.String())

	if priorIncident.FallType != result.FallType {
		t.Error("expected", priorIncident.FallType, "got", result.FallType)
	}

	if priorIncident.AttorneyName != result.AttorneyName {
		t.Error("expected", priorIncident.AttorneyName, "got", result.AttorneyName)
	}

	if priorIncident.LocationID != result.LocationID {
		t.Error("expected", priorIncident.LocationID, "got", result.LocationID)
	}
	if err != nil {
		t.Error("Expected nil ", "got ", err)
	}

	db.Delete(&company)
	db.Delete(&location)
	db.Delete(&priorIncident)
}
