package main

import (
	"fmt"
	"nutrimama/internal/config"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	db, err := config.ConnectDB()
	godotenvErr := godotenv.Load("./.env")
	if godotenvErr != nil {
		panic("Failed to load .env file: " + godotenvErr.Error())
	}

	app := fiber.New()
	config.Bootstrap(&config.BootstrapConfig{
		DB:  db,
		App: app,
	})

	webPort := os.Getenv("APPPORT")
	err = app.Listen(":" + webPort)
	fmt.Printf("Server running on port %s\n", webPort)
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}