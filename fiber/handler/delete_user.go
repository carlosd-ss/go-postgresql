package handler

import (
	"strconv"

	"github.com/carlosd-ss/go-postgresql/fiber/repo"
	"github.com/gofiber/fiber"
)

func (s *Server) Deleteuser(c *fiber.Ctx) {
	params := c.Params("id")
	id, err := strconv.Atoi(params)

	if err != nil {
		c.Status(500).JSON("Não foi possivel converter o id")
	}
	err = repo.DeleteUser(s.Db, int64(id))
	if err != nil {
		c.Status(400).JSON(err)

	}
	c.Status(200)
}
