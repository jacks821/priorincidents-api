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

func TestDeletePriorIncident(t *testing.T) {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)

	if err != nil {
		log.Fatal(err)
		panic("failed to connect to database")
	}
	defer db.Close()

	name := "Meijer"
	company := crud.CreateCompany(name)

	streetNumber := "4123"
	street := "Gibson Rd."
	city := "Henderson"
	state := "Nevada"
	zipCode := "89012"
	storeNumber := "1234"

	location := crud.CreateLocation(streetNumber, street, city, state, zipCode, storeNumber, company.CommonModelFields.ID.String())

	date := time.Now()
	createdDate := date.Format("2006-01-02")
	fallType := "Slip"
	attorneyName := "Charlie Jackson"

	priorIncident := crud.CreatePriorIncident(createdDate, fallType, attorneyName, location.CommonModelFields.ID.String())

	crud.DeletePriorIncident(priorIncident.ID.String())
	crud.DeleteLocation(location.ID.String())

	crud.DeleteCompany(company.ID.String())
}
