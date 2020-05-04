package database

import (
	"github.com/satori/go.uuid"
)

//Report - The Representation of a Report in the database
type Report struct {
	CommonModelFields
	Author          string    `json:"author"`
	Issue           string    `json:"issue"`
	PriorIncidentID uuid.UUID `gorm:"type:uuid;" json:"prior_incident_id"`
	LocationID      uuid.UUID `gorm:"type:uuid;" json:"location_id"`
}
