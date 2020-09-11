package router

import (
	"github.com/carlosd-ss/go-postgresql/fiber/handler"
	"github.com/gofiber/fiber"
)

func SetupRoutes(s *handler.Server, app *fiber.App) {
	app.Delete("/api/v1/user/:id", s.Deleteuser)
	app.Get("/api/v1/user/:id", s.Getuser)
	app.Get("/api/v1/user", s.Getusers)
	app.Post("/api/v1/user", s.Newuser)
}
