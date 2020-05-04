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

func TestCreateCompany(t *testing.T) {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)

	if err != nil {
		log.Fatal(err)
		panic("failed to connect to database")
	}
	defer db.Close()

	name := "Meijer"
	company, err := crud.CreateCompany(name)

	if company.Name != "Meijer" {
		t.Error("expected", name, "got", company.Name)
	}

	if err != nil {
		t.Error("Expected nil ", "got ", err)
	}

	db.Delete(&company)
}
