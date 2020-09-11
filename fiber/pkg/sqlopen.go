package pkg

import (
	"database/sql"
	"fmt"
	"log"
	_ "os"

	_ "github.com/joho/godotenv"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=002212 dbname=user sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// load .env file
	//err := godotenv.Load("db.env")

	//if err != nil {
	//log.Printf("\n%v", err)
	//return nil
	//}

	// Open the connection
	//db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	//if err != nil {
	//log.Printf("\n%v", err)
	//return nil
	//}

	// check the connection
	err = db.Ping()

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}
