package repo

import (
	"database/sql"
	"log"
)

func DeleteUser(db *sql.DB, id int64) error {
	sqlStatement := `DELETE FROM users WHERE userid=$1`

	_, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Println(err)
	}
	return err
}
