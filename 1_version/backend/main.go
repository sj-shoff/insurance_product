package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

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
	dsn := "root:Aesaj2025@tcp(127.0.0.1:3306)/insurance_product?charset=utf8mb4&parseTime=True&loc=Local"
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

func UpdateParameter(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var parameter Parameter
		json.NewDecoder(r.Body).Decode(&parameter)
		db.Save(&parameter)
		json.NewEncoder(w).Encode(&parameter)
	}
}

func DeleteParameter(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.ParseUint(vars["id"], 10, 64)
		db.Delete(&Parameter{}, id)
		json.NewEncoder(w).Encode("Parameter deleted")
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

func UpdateRelationship(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var relationship Relationship
		json.NewDecoder(r.Body).Decode(&relationship)
		db.Save(&relationship)
		json.NewEncoder(w).Encode(&relationship)
	}
}

func DeleteRelationship(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.ParseUint(vars["id"], 10, 64)
		db.Delete(&Relationship{}, id)
		json.NewEncoder(w).Encode("Relationship deleted")
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

func UpdatePartner(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var partner Partner
		json.NewDecoder(r.Body).Decode(&partner)
		db.Save(&partner)
		json.NewEncoder(w).Encode(&partner)
	}
}

func DeletePartner(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.ParseUint(vars["id"], 10, 64)
		db.Delete(&Partner{}, id)
		json.NewEncoder(w).Encode("Partner deleted")
	}
}

func main() {
	router := mux.NewRouter()
	db := InitDB()

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/parameters", CreateParameter(db)).Methods("POST")
	apiRouter.HandleFunc("/parameters", GetParameters(db)).Methods("GET")
	apiRouter.HandleFunc("/parameters", UpdateParameter(db)).Methods("PUT")
	apiRouter.HandleFunc("/parameters/{id}", DeleteParameter(db)).Methods("DELETE")

	apiRouter.HandleFunc("/relationships", CreateRelationship(db)).Methods("POST")
	apiRouter.HandleFunc("/relationships", GetRelationships(db)).Methods("GET")
	apiRouter.HandleFunc("/relationships", UpdateRelationship(db)).Methods("PUT")
	apiRouter.HandleFunc("/relationships/{id}", DeleteRelationship(db)).Methods("DELETE")

	apiRouter.HandleFunc("/partners", CreatePartner(db)).Methods("POST")
	apiRouter.HandleFunc("/partners", GetPartners(db)).Methods("GET")
	apiRouter.HandleFunc("/partners", UpdatePartner(db)).Methods("PUT")
	apiRouter.HandleFunc("/partners/{id}", DeletePartner(db)).Methods("DELETE")

	// Обработка статических файлов
	staticDir := "/home/sj_shoff/insurance_product/first_version/frontend"
	fs := http.FileServer(http.Dir(staticDir))
	router.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(staticDir, r.URL.Path)
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fs.ServeHTTP(w, r)
	}))

	log.Println("Server is running on http://localhost:8000")
	http.ListenAndServe(":8000", router)
}
