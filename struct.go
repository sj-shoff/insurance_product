package entity

type Person_data struct {
	ID           uint     `gorm:"primaryKey"`
	Full_name    string   `gorm:"not null"`
	DateOfBirth  string   `gorm:"not null"`
	Mobile_phone uint     `gorm:"not null"`
	Pasport      string   `gorm:"not null"`
	Pasport_data []string `gorm:"many"`
	Licence      []string `gorm:"many"`
	Registration string   `gorm:"many"`
	Polis        uint     `gorm:"not null"`
	TimeOfUsing  uint     `gorm:"not null"`
}

type Car struct {
	ID          uint          `gorm:"primaryKey"`
	Based_info  []Person_data `gorm:"many"`
	Mark        string        `gorm:"not null"`
	Model       string        `gorm:"not null"`
	Year        uint          `gorm:"not null"`
	Reg_Number  uint          `gorm:"not null"`
	Category    string        `gorm:"not null"`
	Vin         uint          `gorm:"not null"`
	Sts_reg     uint          `gorm:"not null"`
	Engine_info []string      `gorm:"many"`
}

type Health struct {
	ID           uint     `gorm:"primaryKey"`
	Full_name    string   `gorm:"not null"`
	DateOfBirth  string   `gorm:"not null"`
	Mobile_phone uint     `gorm:"not null"`
	Pasport      string   `gorm:"not null"`
	Pasport_data []string `gorm:"many"`
	Licence      []string `gorm:"many"`
	Registration string   `gorm:"many"`
}

type NewProduct struct {
	ProductName   string            `gorm:"not null"`
	Product_param map[string]string `gorn:"not null"`
}

type AllUserProducts struct {
	Products []NewProduct `json:"products"`
}

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
