package database

import (
	"os"

	models "github.com/jacks821/priorincidents-api/database/models"
	"github.com/jinzhu/gorm"
)

//DeleteCompany takes the arguments to delete and soft deletes it from the database.
//Returns an error.
func DeleteCompany(id string) error {
	//s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		return err
	}
	defer db.Close()

	db.Where("id = ?", id).Delete(&models.Company{})
	return nil
}
