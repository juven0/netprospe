package main

import (
	"log"

	"netprospe/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error: .env file not found")
	}

	viewEngine := html.New("./template", ".html")
	 viewEngine.Reload(true)
	// viewEngine.AddFuncMap(template.FuncMap{})

	app := fiber.New(fiber.Config{
		Views: viewEngine,
	})

	app.Static("/static/", "./static")

	app.Get("/", handler.ListeAgent)

	log.Fatal(app.Listen(":3555"))
}
