package main

// func SeedData(db *gorm.DB) {
// 	parameters := []Parameter{
// 		{Name: "Coverage Amount", Type: "number", DefaultValue: "10000"},
// 		{Name: "Policy Term", Type: "number", DefaultValue: "12"},
// 		{Name: "Premium", Type: "number", DefaultValue: "200"},
// 	}
// 	db.Create(&parameters)

// 	relationships := []Relationship{
// 		{Type: "one-to-many", Parameters: parameters[:2]},
// 		{Type: "many-to-one", Parameters: parameters[1:]},
// 	}
// 	db.Create(&relationships)

// 	partners := []Partner{
// 		{Name: "Partner A", Parameters: parameters[:2]},
// 		{Name: "Partner B", Parameters: parameters[1:]},
// 	}
// 	db.Create(&partners)
// }
// SeedData(db)
