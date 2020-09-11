package repo

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/carlosd-ss/go-postgresql/fiber/models/user"
)

func UpdateUser(db *sql.DB, id int64, user user.User) (int64, error) {
	sqlStatement := `UPDATE users SET name=$2, location=$3, age=$4 WHERE userid=$1`

	res, err := db.Exec(sqlStatement, id, user.Name, user.Location, user.Age)

	if err != nil {
		log.Println(err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	if err != nil {
		return rowsAffected, err
	}
	return rowsAffected, nil
}
