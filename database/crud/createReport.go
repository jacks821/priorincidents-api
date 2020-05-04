package database

import (
	"log"

	models "../models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"os"
)

//CreateReport takes the arguments to make a Report and writes it to the database.
//Returns a Report
func CreateReport(author string, issue string, id string, reportType string) models.Report {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", "priorincidents", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)

	if err != nil {
		log.Fatal(err)
		panic("failed to connect to database")
	}
	defer db.Close()

	var report models.Report

	switch reportType {
	case "prior":
		priorID, err := uuid.FromString(id)
		if err != nil {
			log.Fatal(err)
		}
		report.PriorIncidentID = priorID
	case "location":
		locationID, err := uuid.FromString(id)
		if err != nil {
			log.Fatal(err)
		}
		report.LocationID = locationID
	default:
		break
	}

	report.Author = author
	report.Issue = issue

	db.NewRecord(report)
	db.Create(&report)
	return report
}
