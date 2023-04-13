package database

import (
    "fmt"
    "mygram/models"
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
    "os"
)

var (
    host     string
    user     string
    password string
    dbPort   string
    dbname   string
    db       *gorm.DB
    err      error
)

func init() {
    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Set database configuration from environment variables
    host = os.Getenv("DB_HOST")
    user = os.Getenv("DB_USER")
    password = os.Getenv("DB_PASSWORD")
    dbPort = os.Getenv("DB_PORT")
    dbname = os.Getenv("DB_NAME")
}

func StartDB() {
    config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
    dsn := config
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Database connected successfully")
    db.Debug().AutoMigrate(&models.User{}, &models.Comment{}, &models.Photo{}, &models.SocialMedia{})
}

func GetDB() *gorm.DB {
    return db
}