package database

import (
	"github.com/satori/go.uuid"
)

//Location - The Representation of a Location in the database
type Location struct {
	CommonModelFields
	StreetNumber   string          `gorm:"not null" json:"street_number"`
	Street         string          `gorm:"not null" json:"street"`
	City           string          `gorm:"not null" json:"city"`
	State          string          `gorm:"not null" json:"state"`
	ZipCode        string          `gorm:"not null" json:"zip_code"`
	StoreNumber    string          `gorm:"not null" json:"store_number"`
	PriorIncidents []PriorIncident `json:"prior_incidents"`
	CompanyID      uuid.UUID       `gorm:"type:uuid;" json:"company_id"`
	Reports        []Report        `json:"reports"`
}
