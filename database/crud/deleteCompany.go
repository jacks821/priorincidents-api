package database

import (
	models "../models"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

//DeleteCompany takes the arguments to delete and soft deletes it from the database.
//Returns an error.
func DeleteCompany(id string) error {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)

	if err != nil {
		return err
	}
	defer db.Close()

	db.Where("id = ?", id).Delete(&models.Company{})
	return nil
}
