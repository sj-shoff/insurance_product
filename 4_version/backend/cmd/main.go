package main

import (
	"database/sql"
	server "helloapp/4_version/backend/internal"
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const (
	PORT      = 3000
	DATA_FILE = "data.json"
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
	// Убедитесь, что путь к шаблонам правильный
	r.LoadHTMLGlob("4_version/reg/templates/*")

	r.GET("/", server.ShowHomePage)
	r.GET("/register", server.ShowRegistrationForm)
	r.POST("/register", server.RegisterUser)

	// Маршрут для сохранения данных о продукте
	r.POST("/save", server.SaveProductHandler(db))

	//r.GET("/login", server.ShowLoginForm)
	r.POST("/login", server.Login())

	// Запуск сервера
	if err := r.Run(":3000"); err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}

	log.Println("Сервер запущен на http://localhost:3000")
}

//запускать через  go run 4_version/backend/cmd/main.go
