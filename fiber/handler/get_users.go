package handler

import (
	"log"

	"github.com/carlosd-ss/go-postgresql/fiber/repo"
	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
)

func Getusers(c *fiber.Ctx) {
	users, err := repo.GetAllUsers()

	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}
	c.JSON(users)

}
