package database

import (
	"os"

	models "github.com/jacks821/priorincidents-api/database/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//GetCompanyByLocation finds a Company by the ID of one of its Locations in the database.
//Returns a Company.
func GetCompanyByLocation(locationID string) (models.Company, error) {
	//s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	var company models.Company
	if err != nil {
		return company, err
	}
	defer db.Close()

	locID, err := uuid.FromString(locationID)
	if err != nil {
		return company, err
	}

	var location models.Location
	db.Where("id = ?", locID).Find(&location)
	db.Where("id = ?", location.CompanyID).Find(&company)
	return company, nil
}
