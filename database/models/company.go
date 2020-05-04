package database

//Company - The Representation of a Company in the database
type Company struct {
	CommonModelFields
	Name      string     `gorm:"unique;not null" json:"name"`
	Locations []Location `json:"locations"`
}
