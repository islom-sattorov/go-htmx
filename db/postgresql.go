package db

import (
	"go_role/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(dsn string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}

// Add it to the main
// // dns := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable",
// 	os.Getenv("DB_HOST"),
// 	os.Getenv("DB_USER"),
// 	os.Getenv("DB_PASSWORD"),
// 	os.Getenv("DB_NAME"),
// )
// database, err := db.NewDatabase(dns)
// if err != nil {
// 	log.Fatal(err)
// }

// user := models.User{
// 	Name:  "John Doe",
// 	Email: "john@example.com",
// 	Role:  "admin",
// }

// result := database.DB.Create(&user)
// if result.Error != nil {
// 	log.Fatal(result.Error)
// }
