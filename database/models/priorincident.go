package database

import (
	"github.com/satori/go.uuid"
	"time"
)

//PriorIncident - The Representation of a PriorIncident in the database
type PriorIncident struct {
	CommonModelFields
	Date         time.Time `json:"date"`
	FallType     string    `json:"fall_type"`
	AttorneyName string    `json:"attorney_name"`
	LocationID   uuid.UUID `gorm:"type:uuid;" json:"location_id"`
	Reports      []Report  `json:"reports"`
}
