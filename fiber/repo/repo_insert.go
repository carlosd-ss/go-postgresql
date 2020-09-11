package repo

import (
	"database/sql"

	"github.com/carlosd-ss/go-postgresql/fiber/models/user"
)

func InsertUser(db *sql.DB, user user.User) (int64, error) {
	sqlStatement := `INSERT INTO users (name, location, age) VALUES ($1, $2, $3) RETURNING userid`

	var id int64

	err := db.QueryRow(sqlStatement, user.Name, user.Location, user.Age).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}
