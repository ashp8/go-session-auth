package connection

import (
	"fmt"

	"com.ashp8/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func SetupDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=session-auth sslmode=disable")
	DB = db
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	fmt.Println("Database Connection established!")
	return db, nil
}
