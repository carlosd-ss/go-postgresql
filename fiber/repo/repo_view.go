package repo

import (
	"database/sql"

	"github.com/carlosd-ss/go-postgresql/fiber/models/user"
)

func GetUser(db *sql.DB, id int64) (user.User, error) {
	var user user.User

	sqlStatement := `SELECT * FROM users WHERE userid=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

	if err != nil {
		return user, err
	}
	return user, nil
}
