package main

import (
	"database/sql"
	server "helloapp/4_version/backend/internal"
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:Aesaj2025@tcp(localhost)/insurance_product")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = server.LoadDataFromFile(db, filepath.Join("/home/sj_shoff/insurance_product/4_version/frontend", "data.json"))
	if err != nil {
		log.Fatal("Ошибка при чтении данных из файла: ", err)
	}

	r := gin.Default()
	r.LoadHTMLGlob("4_version/reg/templates/*")

	r.GET("/", server.ShowHomePage)
	r.GET("/register", server.ShowRegistrationForm)
	r.POST("/register", server.RegisterUser)
	r.POST("/login", server.Login())

	// Маршрут для сохранения данных о продукте
	r.POST("/save", server.SaveProductHandler(db))

	// Запуск сервера на порту 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}

	log.Println("Сервер запущен на http://localhost:8080")
}

//запускать через  go run 4_version/backend/cmd/main.go
