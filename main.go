package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	crud "github.com/jacks821/priorincidents-api/database/crud"
	migrations "github.com/jacks821/priorincidents-api/migrations"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	migrations.RunMigrations()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/companies", listCompanies).Methods("GET")
	router.HandleFunc("/company", createCompany).Methods("POST")
	router.HandleFunc("/company", createCompany).Methods("OPTIONS")
	router.HandleFunc("/report", createReport).Methods("POST")
	router.HandleFunc("/companies/{id}", getCompany).Methods("GET")
	router.HandleFunc("/company/location={id}", getCompanyByLocation).Methods("GET")
	router.HandleFunc("/companies/location/{id}", getLocation).Methods("GET")
	router.HandleFunc("/company/location", createLocation).Methods("POST")
	router.HandleFunc("/company/location/priorincident", createPriorIncident).Methods("POST")
	router.HandleFunc("/companies/location/priorIncident/{id}", getPriorIncident).Methods("GET")
	router.HandleFunc("/companies/{id}", deleteCompany).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func createCompany(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	fmt.Println("Entered createCompany Handler")
	var result map[string]interface{}
	json.NewDecoder(r.Body).Decode(&result)
	name := fmt.Sprintf("%v", result["name"])
	company, err := crud.CreateCompany(name)
	if err != nil {
		log.Printf("Database error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Successfully created Company: ", company.Name)
}

func createLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered createLocation Handler")
	var result map[string]interface{}
	json.NewDecoder(r.Body).Decode(&result)
	street := fmt.Sprintf("%v", result["street"])
	streetNumber := fmt.Sprintf("%v", result["street_number"])
	city := fmt.Sprintf("%v", result["city"])
	state := fmt.Sprintf("%v", result["state"])
	zipCode := fmt.Sprintf("%v", result["zip_code"])
	storeNumber := fmt.Sprintf("%v", result["store_number"])
	companyID := fmt.Sprintf("%v", result["company_id"])
	_, err := crud.CreateLocation(streetNumber, street, city, state, zipCode, storeNumber, companyID)
	if err != nil {
		log.Printf("Database error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Successfully created Location!")
}

func createPriorIncident(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered createPriorIncident Handler")
	var result map[string]interface{}
	json.NewDecoder(r.Body).Decode(&result)
	date := fmt.Sprintf("%v", result["date"])
	fallType := fmt.Sprintf("%v", result["fall_type"])
	attorneyName := fmt.Sprintf("%v", result["attorney_name"])
	locationID := fmt.Sprintf("%v", result["location_id"])
	_, err := crud.CreatePriorIncident(date, fallType, attorneyName, locationID)
	if err != nil {
		log.Printf("Database error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Successfully created Prior Incident!")
}

func createReport(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered createReport Handler")
	var result map[string]interface{}
	json.NewDecoder(r.Body).Decode(&result)
	author := fmt.Sprintf("%v", result["author"])
	id := fmt.Sprintf("%v", result["id"])
	issue := fmt.Sprintf("%v", result["issue"])
	reportType := fmt.Sprintf("%v", result["report_type"])
	_, err := crud.CreateReport(author, issue, id, reportType)
	if err != nil {
		log.Printf("Database error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Successfully created Report!")
}

func listCompanies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered listcompanies Handler")
	companies, err := crud.ListCompanies()
	if err != nil {
		log.Printf("Database error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(companies)
}

func getCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered getCompany Handler")
	companyID := r.URL.Path[len("companies/")+1:]

	company, err := crud.GetCompany(companyID)
	if err != nil {
		log.Printf("Database error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(company)
}

func getCompanyByLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered getCompanyByLocation Handler")
	locationID := r.URL.Path[len("company/location=")+1:]

	company, err := crud.GetCompanyByLocation(locationID)
	if err != nil {
		log.Printf("Database error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(company)
}

func getLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered getLocation Handler")
	locationID := r.URL.Path[len("companies/location/")+1:]

	location, err := crud.GetLocation(locationID)
	if err != nil {
		log.Printf("Database error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(location)
}

func getPriorIncident(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered getPriorIncident Handler")
	priorIncidentID := r.URL.Path[len("companies/location/priorIncident/")+1:]
	priorIncident, err := crud.GetPriorIncident(priorIncidentID)
	if err != nil {
		log.Printf("Database error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(priorIncident)
}

func deleteCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered deleteCompany handler")
	companyID := r.URL.Path[1:]

	err := crud.DeleteCompany(companyID)
	if err != nil {
		log.Printf("Database error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Successfully deleted Company")
}

func deleteLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered deleteLocation handler")
	locationID := r.URL.Path[1:]

	err := crud.DeleteLocation(locationID)
	if err != nil {
		log.Printf("Database error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Println("Successfully deleted Location")
}

func deletePriorIncident(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered deletePriorIncident handler")
	priorIncidentID := r.URL.Path[1:]

	err := crud.DeletePriorIncident(priorIncidentID)
	if err != nil {
		log.Printf("Database error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Successfully deleted Prior Incident")
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
