package main

import (
	"github.com/carlosd-ss/go-postgresql/fiber/handler"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Delete("/api/v1/user/:id", handler.Deleteuser)
	app.Get("/api/v1/user/:id", handler.Getuser)
	app.Get("/api/v1/user", handler.Getusers)
	app.Post("/api/v1/user", handler.Newuser)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
}
