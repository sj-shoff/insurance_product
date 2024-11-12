# insurance_product for hack
func InitDB() *gorm.DB {
	dsn := "root:Aesaj2025@tcp(127.0.0.1:3306)/insurance product?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&Parameter{}, &Relationship{}, &Partner{})
	return db
}