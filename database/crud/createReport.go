package database

import (
	"fmt"
	"os"

	models "github.com/jacks821/priorincidents-api/database/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//CreateReport takes the arguments to make a Report and writes it to the database.
//Returns a Report and an error
func CreateReport(author string, issue string, id string, reportType string) (models.Report, error) {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	var report models.Report
	if err != nil {
		return report, err
	}
	defer db.Close()

	switch reportType {
	case "prior":
		priorID, err := uuid.FromString(id)
		if err != nil {
			return report, err
		}
		report.PriorIncidentID = priorID
	case "location":
		locationID, err := uuid.FromString(id)
		if err != nil {
			return report, err
		}
		report.LocationID = locationID
	default:
		break
	}

	report.Author = author
	report.Issue = issue

	db.NewRecord(report)
	db.Create(&report)
	return report, nil
}
