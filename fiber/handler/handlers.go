package handler

import (
	"fmt"
	"log"
	"strconv"

	"github.com/carlosd-ss/go-postgresql/fiber/models"
	"github.com/carlosd-ss/go-postgresql/fiber/repo"
	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
)

func Getusers(c *fiber.Ctx) {
}

func Getuser(c *fiber.Ctx) {
}

func Newuser(c *fiber.Ctx) {
	var p models.User
	if err := c.BodyParser(&p); err != nil {
		c.Status(503).Send(err)
		return
	}
	insert := repo.InsertUser(p)
	c.JSON(insert)
}

func Deleteuser(c *fiber.Ctx) {
	params := c.Params("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	deletedRows := repo.DeleteUser(int64(id))
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)
	c.JSON(msg)
}
