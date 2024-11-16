package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	entity "helloapp"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:111@tcp(localhost)/insurance_product")
	if err != nil {
		log.Fatal(err)
	}
}

func CreateUser(user *entity.User) error {
	stmt, err := db.Prepare("INSERT INTO users(username, password) VALUES(?, ?)")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Username, user.Password)
	return err
}

func ShowRegistrationForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func RegisterUser(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBind(&user); err != nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Invalid form data"})
		return
	}

	// Валидация данных
	if user.Username == "" || user.Password == "" {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Username and password are required"})
		return
	}

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// Сохранение пользователя в базе данных
	if err := CreateUser(&user); err != nil {
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{"error": "Failed to create user"})
		return
	}

	c.HTML(http.StatusOK, "register.html", gin.H{"success": "Registration successful"})
}

func ShowHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func AddNewProductPattern(data_entry []byte, curent_user_id uint) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		last := entity.NewProduct{}
		json.Unmarshal(data_entry, &last)

		var AllProducts []byte
		db.QueryRow("SELECT all_products FROM users WHERE id = ?", curent_user_id).Scan(&AllProducts)
		var Products []entity.NewProduct
		json.Unmarshal(AllProducts, &Products)
		Products = append(Products, last)
		AllProducts, err := json.Marshal(Products)
		if err != nil {
			log.Fatal("failed marshal")
		}

		stmt, err := db.Prepare("INSERT INTO users(new_products, all_products) VALUES(?, ?)")
		if err != nil {
			log.Fatal("no db")
		}
		defer stmt.Close()

		_, err = stmt.Exec(data_entry, AllProducts)
		if err != nil {
			log.Fatal("bad db")
		}
	}
}
