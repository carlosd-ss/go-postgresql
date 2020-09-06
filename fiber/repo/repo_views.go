package repo

import (
	"log"

	"github.com/carlosd-ss/go-postgresql/fiber/models/user"
	"github.com/carlosd-ss/go-postgresql/fiber/pkg"
)

func GetAllUsers() ([]user.User, error) {

	db := pkg.CreateConnection()

	defer db.Close()

	var users []user.User

	sqlStatement := `SELECT * FROM users`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user user.User

		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		users = append(users, user)

	}

	return users, nil
}
