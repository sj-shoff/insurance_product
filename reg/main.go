package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:Aesaj2025@tcp(127.0.0.1:3306)/users_reg")
	if err != nil {
		log.Fatal(err)
	}
}

func CreateUser(user *User) error {
	stmt, err := db.Prepare("INSERT INTO users(username, password) VALUES(?, ?)")
	if err != nil {
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
	var user User

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

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", ShowHomePage)
	r.GET("/register", ShowRegistrationForm)
	r.POST("/register", RegisterUser)

	r.Run(":8080")
}
