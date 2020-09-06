package main

import (
	"github.com/carlosd-ss/go-postgresql/fiber/router"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	router.SetupRoutes(app)
	app.Listen(3000)
}
