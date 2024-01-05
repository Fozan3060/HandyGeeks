package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/template/html/v2"
	"server/pkg/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New(fiber.Config{})
	app.Use(cors.New())
	routes.Setup(app)

	log.Println("[+] Server is running on port 3000]")
	app.Listen(":3000")

}
