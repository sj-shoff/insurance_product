package main

import (
	server "helloapp/4_version/backend/internal"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Убедитесь, что путь к шаблонам правильный
	r.LoadHTMLGlob("4_version/reg/templates/*")

	r.GET("/", server.ShowHomePage)
	r.GET("/register", server.ShowRegistrationForm)
	r.POST("/register", server.RegisterUser)

	//r.GET("/login", server.ShowLoginForm)
	r.POST("/login", server.Login())

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}

	log.Println("Сервер запущен на http://localhost:8080")
}

//запускать через  go run 4_version/backend/cmd/main.go
