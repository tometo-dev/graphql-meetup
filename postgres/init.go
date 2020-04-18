package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/tsuki42/graphql-meetup/logging"
	"os"
)

var Connection *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		logging.ERROR.Println("Error loading .env file")
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	URI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName, dbPassword)

	Connection, err = gorm.Open("postgres", URI)

	if err != nil {
		logging.ERROR.Println("Failed to connect to the database")
		panic(err)
	} else {
		logging.INFO.Println("Successfully connected to the database")
	}

	Connection.LogMode(true)
	Connection.SetLogger(logging.DB)
	Connection.DB().SetMaxOpenConns(50)
	Connection.DB().SetMaxIdleConns(10)

	ValidateSchema()

}
