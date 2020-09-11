package main

import (
	"github.com/carlosd-ss/go-postgresql/fiber/handler"
	"github.com/carlosd-ss/go-postgresql/fiber/pkg"
	"github.com/carlosd-ss/go-postgresql/fiber/router"
	"github.com/gofiber/fiber"
)

func main() {
	s := handler.Server{
		Db: pkg.CreateConnection(),
	}
	defer s.Db.Close()

	app := fiber.New()
	router.SetupRoutes(&s, app)
	app.Listen(3000)
}
