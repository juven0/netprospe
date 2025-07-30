package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error: .env file not found")
	}

	viewEngine := html.New("./template", "html")

	app := fiber.New(fiber.Config{
		Views: viewEngine,
	})

	log.Fatal(app.Listen(":3555"))
}
