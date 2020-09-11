package handler

import (
	"database/sql"

	"github.com/gofiber/fiber"
)

type Server struct {
	App *fiber.App
	Db  *sql.DB
}
