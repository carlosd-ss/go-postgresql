package handler

import (
	"fmt"
	"strconv"

	"github.com/carlosd-ss/go-postgresql/fiber/repo"
	"github.com/gofiber/fiber"
)

func Getuser(c *fiber.Ctx) {

	params := c.Params("id")
	id, err := strconv.Atoi(params)

	if err != nil {
		c.Status(400).JSON(err)
	}
	Row, err := repo.GetUser(int64(id))
	if err != nil {
		c.Status(400).JSON(err)

	}

	msg := fmt.Sprintf("%v", Row)
	c.Status(200).JSON(msg)
}
