package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/carlosd-ss/go-postgresql/gorila-mux/router"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	r := router.Router()

	fmt.Println("Server rodando na porta 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
