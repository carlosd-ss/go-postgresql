package repo

import (
	"database/sql"

	"github.com/carlosd-ss/go-postgresql/fiber/models/user"
)

func GetAllUsers(db *sql.DB) ([]user.User, error) {

	var users []user.User

	sqlStatement := `SELECT * FROM users`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user user.User

		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

		if err != nil {
			return nil, err
		}

		users = append(users, user)

	}

	return users, nil
}
