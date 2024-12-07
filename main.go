package main

import (
	"divyanshu050303/insta_backend/database"
	"divyanshu050303/insta_backend/models"
	"divyanshu050303/insta_backend/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	fmt.Println("Database Config:", config)
	db, err := database.NewConnection(config)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		log.Fatal("Could Not load datebase")
	}
	fmt.Println("Database Connection Established")
	err = models.Migrate(db)
	if err != nil {
		log.Fatal("Count not migrate the databse")
	}
	app := fiber.New()
	routes.SetUpUserRoutes(app, db)
	app.Listen(":8000")

}
