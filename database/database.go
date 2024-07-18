package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/juniorteye/devCamp/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database instance
type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// Connect function
func Connect() {
	// Load environment variables from .env file
	err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// Fetch and validate environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Debugging prints
	fmt.Println("DB_HOST:", dbHost)
	fmt.Println("DB_USER:", dbUser)
	fmt.Println("DB_PASSWORD:", dbPassword)
	fmt.Println("DB_NAME:", dbName)
	fmt.Println("DB_PORT:", dbPort)

	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbPort == "" {
		log.Fatal("Database environment variables are not set")
	}

	port, err := strconv.ParseUint(dbPort, 10, 32)
	if err != nil {
		log.Fatal("Error parsing DB_PORT to uint: ", err)
	}

	// Construct DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPassword, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&model.User{}, &model.Review{})
	DB = Dbinstance{
		Db: db,
	}
}
