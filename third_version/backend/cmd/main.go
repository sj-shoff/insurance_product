package main

import (
	server "helloapp/third_version/backend/internal"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	db := server.InitDB()

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/parameters", server.CreateParameter(db)).Methods("POST")
	apiRouter.HandleFunc("/parameters", server.GetParameters(db)).Methods("GET")
	apiRouter.HandleFunc("/parameters", server.UpdateParameter(db)).Methods("PUT")
	apiRouter.HandleFunc("/parameters/{id}", server.DeleteParameter(db)).Methods("DELETE")

	apiRouter.HandleFunc("/relationships", server.CreateRelationship(db)).Methods("POST")
	apiRouter.HandleFunc("/relationships", server.GetRelationships(db)).Methods("GET")
	apiRouter.HandleFunc("/relationships", server.UpdateRelationship(db)).Methods("PUT")
	apiRouter.HandleFunc("/relationships/{id}", server.DeleteRelationship(db)).Methods("DELETE")

	apiRouter.HandleFunc("/partners", server.CreatePartner(db)).Methods("POST")
	apiRouter.HandleFunc("/partners", server.GetPartners(db)).Methods("GET")
	apiRouter.HandleFunc("/partners", server.UpdatePartner(db)).Methods("PUT")
	apiRouter.HandleFunc("/partners/{id}", server.DeletePartner(db)).Methods("DELETE")

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
