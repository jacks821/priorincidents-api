package database

import (
	models "../models"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

//DeleteLocation takes the arguments to delete a Location and soft deletes it from the database.
func DeleteLocation(id string) {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)

	if err != nil {
		log.Fatal(err)
		panic("failed to connect to database")
	}
	defer db.Close()

	db.Where("id = ?", id).Delete(&models.Location{})
}
