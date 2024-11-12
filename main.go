package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Parameter struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	Type         string `gorm:"not null"`
	DefaultValue string
	DictionaryID uint
}

type Relationship struct {
	ID         uint        `gorm:"primaryKey"`
	Type       string      `gorm:"not null"`
	Parameters []Parameter `gorm:"many2many:relationship_parameters;"`
}

type Partner struct {
	ID         uint        `gorm:"primaryKey"`
	Name       string      `gorm:"not null"`
	Parameters []Parameter `gorm:"many2many:partner_parameters;"`
}

func InitDB() *gorm.DB {
	dsn := "root:Aesaj2025@tcp(127.0.0.1:3306)/insurance product?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&Parameter{}, &Relationship{}, &Partner{})
	return db
}

func CreateParameter(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var parameter Parameter
		json.NewDecoder(r.Body).Decode(&parameter)
		db.Create(&parameter)
		json.NewEncoder(w).Encode(&parameter)
	}
}

func GetParameters(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var parameters []Parameter
		db.Find(&parameters)
		json.NewEncoder(w).Encode(&parameters)
	}
}

func CreateRelationship(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var relationship Relationship
		json.NewDecoder(r.Body).Decode(&relationship)
		db.Create(&relationship)
		json.NewEncoder(w).Encode(&relationship)
	}
}

func GetRelationships(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var relationships []Relationship
		db.Preload("Parameters").Find(&relationships)
		json.NewEncoder(w).Encode(&relationships)
	}
}

func CreatePartner(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var partner Partner
		json.NewDecoder(r.Body).Decode(&partner)
		db.Create(&partner)
		json.NewEncoder(w).Encode(&partner)
	}
}

func GetPartners(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var partners []Partner
		db.Preload("Parameters").Find(&partners)
		json.NewEncoder(w).Encode(&partners)
	}
}

func main() {
	router := mux.NewRouter()
	db := InitDB()

	router.HandleFunc("/parameters", CreateParameter(db)).Methods("POST")
	router.HandleFunc("/parameters", GetParameters(db)).Methods("GET")
	router.HandleFunc("/relationships", CreateRelationship(db)).Methods("POST")
	router.HandleFunc("/relationships", GetRelationships(db)).Methods("GET")
	router.HandleFunc("/partners", CreatePartner(db)).Methods("POST")
	router.HandleFunc("/partners", GetPartners(db)).Methods("GET")

	http.ListenAndServe(":8000", router)
}
