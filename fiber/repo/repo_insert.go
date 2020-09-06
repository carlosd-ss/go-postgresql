package repo

import (
	"github.com/carlosd-ss/go-postgresql/fiber/models/user"
	"github.com/carlosd-ss/go-postgresql/fiber/pkg"
)

func InsertUser(user user.User) (int64, error) {

	db := pkg.CreateConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO users (name, location, age) VALUES ($1, $2, $3) RETURNING userid`

	var id int64

	err := db.QueryRow(sqlStatement, user.Name, user.Location, user.Age).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}
