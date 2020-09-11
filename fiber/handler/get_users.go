package handler

import (
	"github.com/carlosd-ss/go-postgresql/fiber/repo"
	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
)

func (s *Server) Getusers(c *fiber.Ctx) {
	users, err := repo.GetAllUsers(s.Db)

	if err != nil {
		c.Status(400).JSON(err)
	}

	c.JSON(users)

}
