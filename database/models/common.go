package database

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //Needed for database
	uuid "github.com/satori/go.uuid"
)

//CommonModelFields are the fields which will exist on all models in the database.
type CommonModelFields struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//BeforeCreate runs before the entry of any entry to the database, this function will be run, giving a UUID to the model.
func (base *CommonModelFields) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}
