package main

import (
	"divyanshu050303/insta_backend/database"
	"divyanshu050303/insta_backend/models"
	"divyanshu050303/insta_backend/routes"
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
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	db, err := database.NewConnection(config)
	if err != nil {
		log.Fatal("Cound not load the database")
	}
	err = models.Migrate(db)
	if err != nil {
		log.Fatal("Cound not migrate the database")
	}
	app := fiber.New()
	routes.SetUpUserRoutes(app, db)
	app.Listen(":8000")

}
