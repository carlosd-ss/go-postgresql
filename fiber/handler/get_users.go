package handler

import (
	"github.com/carlosd-ss/go-postgresql/fiber/repo"
	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
)

func Getusers(c *fiber.Ctx) {
	users, err := repo.GetAllUsers()

	if err != nil {
		c.Status(400).JSON(err)
	}

	c.JSON(users)

}
