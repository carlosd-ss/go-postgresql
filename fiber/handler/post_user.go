package handler

import (
	"github.com/carlosd-ss/go-postgresql/fiber/models/merrors"
	"github.com/carlosd-ss/go-postgresql/fiber/models/user"
	"github.com/carlosd-ss/go-postgresql/fiber/repo"
	"github.com/gofiber/fiber"
)

func Newuser(c *fiber.Ctx) {
	var p user.User
	var msgE merrors.Errors

	if err := c.BodyParser(&p); err != nil {
		c.Status(503).Send(err)
		return
	}
	if p.Age <= 0 {
		msgE.Msg = "Campo Obrigatorio"
		c.Status(400).JSON(msgE)
		return
	}

	id, err := repo.InsertUser(p)
	if err != nil {
		c.Status(400).JSON(err)
	}
	c.Status(200).JSON(id)
}
