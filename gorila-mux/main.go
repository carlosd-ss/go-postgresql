package main

import (
	"fmt"
	"github.com/carlosd-ss/go-cruds/examples/gorila-mux/router"
	"log"
	"net/http"
	"os"
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