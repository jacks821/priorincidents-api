package migrations

import (
	models "../database/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //Required for database to work
	"log"
	"os"
)

//RunMigrations runs the database Migrations and prints success to console.
func RunMigrations() {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.AutoMigrate(&models.Company{}, &models.Location{}, &models.PriorIncident{}, &models.Report{})
	fmt.Println("Successful Migration!")
}
