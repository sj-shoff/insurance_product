package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	entity "helloapp"
	"io/ioutil"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var store = sessions.NewCookieStore([]byte("secret-key"))

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entity.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Выполнение SQL-запроса для поиска пользователя
		var foundUser entity.User
		query := "SELECT * FROM users WHERE username = ? AND password = ?"
		row := db.QueryRow(query, user.Username, user.Password)

		// Сканирование результата в foundUser
		err := row.Scan(&foundUser.ID, &foundUser.Username, &foundUser.Password) // Убедитесь, что поля соответствуют вашей структуре
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			}
			return
		}

		// Здесь вы можете создать сессию или JWT-токен

		session, err := store.Get(c.Request, "session-name")
		if err != nil {
			c.JSON(500, gin.H{"error": "Could not get session"})
			return
		}
		userID := foundUser.ID
		// Сохраняем user_id в сессии
		session.Values["user_id"] = userID
		if err := session.Save(c.Request, c.Writer); err != nil {
			c.JSON(500, gin.H{"error": "Could not save session"})
			return
		}

		c.JSON(200, gin.H{"message": "Logged in successfully", "user_id": userID})
	}
}

func saveProductHandler(c *gin.Context) {
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var err error
	db, err = sql.Open("mysql", "root:Aesaj2025@tcp(127.0.0.1:3306)/new_product")
	if err != nil {
		log.Fatal(err)
	}

	// Здесь вы можете добавить логику для сохранения данных в базу данных или другое хранилище
	fmt.Printf("Received product: %+v\n", product)

	c.JSON(http.StatusCreated, gin.H{"message": "Product saved successfully"})
}

func GetIdOfSession(c *gin.Context) uint {
	session, err := store.Get(c.Request, "session-name")
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not get session"})
		return 0
	}

	userID, ok := session.Values["user_id"].(int) // Приводим к правильному типу
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return 0
	}

	// Здесь вы можете использовать userID для чего-либо
	c.JSON(200, gin.H{"message": "Protected data", "user_id": userID})
	return uint(userID)
}

func init() {
	var err error
	db, err = sql.Open("mysql", "root:Aesaj2025@tcp(localhost)/insurance_product")
	if err != nil {
		log.Fatal(err)
	}
}

func CreateUser(user *entity.User) error {
	stmt, err := db.Prepare("INSERT INTO users(username, password) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	fmt.Println("oki")
	_, err = stmt.Exec(user.Username, user.Password)
	fmt.Println("NOToki")
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
		fmt.Println(err)
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{"error": "Failed to create user"})
		return
	}

	c.HTML(http.StatusOK, "register.html", gin.H{"success": "Registration successful"})
}

func ShowHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func AddNewProductPattern(data_entry []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		curent_user_id := GetIdOfSession(ctx)

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

const (
	DATA_FILE = "data.json"
)

func SaveProductHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product entity.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Преобразуем данные в JSON
		productJSON, err := json.Marshal(product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
			return
		}

		// Сохраняем данные в базу данных
		query := `
			INSERT INTO products (name, start_date, end_date, update_date, version_description, series_prefix, series_postfix, number_prefix, number_postfix, numerator, custom_number_method, individual_parameters, cost_formula)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`
		_, err = db.Exec(query, product.Name, product.StartDate, product.EndDate, product.UpdateDate, product.VersionDescription, product.SeriesPrefix, product.SeriesPostfix, product.NumberPrefix, product.NumberPostfix, product.Numerator, product.CustomNumberMethod, productJSON, product.CostFormula)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Data saved successfully"})
	}
}

func LoadDataFromFile(db *sql.DB, filePath string) error {
	// Чтение данных из файла
	data, err := ioutil.ReadFile(DATA_FILE)
	if err != nil {
		return fmt.Errorf("ошибка при чтении файла: %v", err)
	}

	// Декодирование JSON-данных
	var product entity.Product
	err = json.Unmarshal(data, &product)
	if err != nil {
		return fmt.Errorf("ошибка при декодировании JSON: %v", err)
	}

	// Преобразуем данные в JSON
	productJSON, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("ошибка при маршалинге данных: %v", err)
	}

	// Сохраняем данные в базу данных
	query := `
		INSERT INTO products (name, start_date, end_date, update_date, version_description, series_prefix, series_postfix, number_prefix, number_postfix, numerator, custom_number_method, individual_parameters, cost_formula)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err = db.Exec(query, product.Name, product.StartDate, product.EndDate, product.UpdateDate, product.VersionDescription, product.SeriesPrefix, product.SeriesPostfix, product.NumberPrefix, product.NumberPostfix, product.Numerator, product.CustomNumberMethod, productJSON, product.CostFormula)
	if err != nil {
		return fmt.Errorf("ошибка при сохранении данных в базу данных: %v", err)
	}

	return nil
}
