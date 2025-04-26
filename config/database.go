package config

import (
	"fmt"
	"log"

	"github.com/nikhil/blogging-platform-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=localhost user=postgres password=admin dbname=blogging port=5432 sslmode=disable"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    // Auto-migrate Post model
    err = database.AutoMigrate(&models.Post{})
    if err != nil {
        log.Fatal("Migration failed: ", err)
    }

    DB = database
    fmt.Println("Database connected successfully.")
}
