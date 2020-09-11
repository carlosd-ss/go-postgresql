package handler

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/carlosd-ss/go-postgresql/fiber/repo"
	"github.com/gofiber/fiber"
)

type Server struct {
	app *fiber.App
	db  *sql.DB
}

func Deleteuser(c *fiber.Ctx) {
	params := c.Params("id")
	id, err := strconv.Atoi(params)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	deletedRows, err := repo.DeleteUser(int64(id))
	if err != nil {
		c.Status(400).JSON(err)

	}

	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)
	c.Status(200).JSON(msg)
}
