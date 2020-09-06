package handler

import (
	"fmt"
	"log"
	"strconv"

	"github.com/carlosd-ss/go-postgresql/fiber/repo"
	"github.com/gofiber/fiber"
)

func Getuser(c *fiber.Ctx) {

	params := c.Params("id")
	id, err := strconv.Atoi(params)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	Row, err := repo.GetUser(int64(id))
	if err != nil {
		log.Fatalf("Unable to view.  %v", err)

	}

	msg := fmt.Sprintf("%v", Row)
	c.Status(200).JSON(msg)
}
