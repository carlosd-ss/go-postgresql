package repo

import (
	"github.com/carlosd-ss/go-postgresql/fiber/models/user"
	"github.com/carlosd-ss/go-postgresql/fiber/pkg"
)

func GetUser(id int64) (user.User, error) {

	db := pkg.CreateConnection()

	defer db.Close()

	var user user.User

	sqlStatement := `SELECT * FROM users WHERE userid=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

	if err != nil {
		return user, err
	}
	return user, nil
}
