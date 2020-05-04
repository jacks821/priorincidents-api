module github.com/jacks821/priorincidents-api

go 1.14

require (
    "github.com/jinzhu/gorm"
    "./database/crud"
    "github.com/jinzhu/gorm/dialects/postgres"
    "./database/models"
    "./migrations"
    "github.com/gorilla/mux"
)
