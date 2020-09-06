package handler

import (
	"fmt"
	"log"
	"strconv"

	"github.com/carlosd-ss/go-postgresql/fiber/repo"
	"github.com/gofiber/fiber"
)

func Deleteuser(c *fiber.Ctx) {
	params := c.Params("id")
	id, err := strconv.Atoi(params)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	deletedRows, err := repo.DeleteUser(int64(id))
	if err != nil {
		log.Fatalf("Unable to delete.  %v", err)

	}

	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)
	c.Status(200).JSON(msg)
}
