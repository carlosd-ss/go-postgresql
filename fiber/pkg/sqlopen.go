package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func CreateConnection() *sql.DB {
	// load .env file
	err := godotenv.Load("db.env")

	if err != nil {
		log.Printf("\n%v", err)
		return nil
	}

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		log.Printf("\n%v", err)
		return nil
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}
