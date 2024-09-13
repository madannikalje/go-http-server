package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var DB_HOST string
var DB_PORT string
var DB_USER string
var DB_PASSWORD string
var API_KEY string
var DB_NAME string
var JWT_SECRET string

func Init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
		panic(err)
	}
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	API_KEY = os.Getenv("API_KEY")
	DB_NAME = os.Getenv("DB_NAME")
	JWT_SECRET = os.Getenv("JWT_SECRET")
}
