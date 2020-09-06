package router

import (
	"github.com/carlosd-ss/go-postgresql/fiber/handler"
	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {
	app.Delete("/api/v1/user/:id", handler.Deleteuser)
	app.Get("/api/v1/user/:id", handler.Getuser)
	app.Get("/api/v1/user", handler.Getusers)
	app.Post("/api/v1/user", handler.Newuser)
}
