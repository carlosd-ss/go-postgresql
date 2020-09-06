package repo

import (
	"fmt"
	"log"

	"github.com/carlosd-ss/go-postgresql/fiber/models/user"
	"github.com/carlosd-ss/go-postgresql/fiber/pkg"
)

func DeleteUser(id int64) (user.User, error) {
	db := pkg.CreateConnection()

	defer db.Close()
	var user user.User
	sqlStatement := `DELETE FROM users WHERE userid=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return user, err
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return user, err
}
