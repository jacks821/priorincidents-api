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
	models "../models"
)

func TestCreatePriorIncident(t *testing.T) {
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

	crud.CreateCompany("Meijer")
	var company models.Company

	db.Where("name = ?", "Meijer").First(&company)

	location, _ := crud.CreateLocation(streetNumber, street, city, state, zipCode, storeNumber, company.CommonModelFields.ID.String())
	date := time.Now()
	createdDate := date.Format("2006-01-02")
	fallType := "Slip"
	attorneyName := "Charlie Jackson"

	priorIncident, err := crud.CreatePriorIncident(createdDate, fallType, attorneyName, location.CommonModelFields.ID.String())

	if priorIncident.FallType != fallType {
		t.Error("expected", fallType, "got", priorIncident.FallType)
	}

	if priorIncident.AttorneyName != attorneyName {
		t.Error("expected", attorneyName, "got", priorIncident.AttorneyName)
	}

	if priorIncident.LocationID != location.CommonModelFields.ID {
		t.Error("expected", location.ID, "got", priorIncident.LocationID)
	}
	if err != nil {
		t.Error("Expected nil ", "got ", err)
	}

	db.Delete(&company)
	db.Delete(&location)
	db.Delete(&priorIncident)
}
