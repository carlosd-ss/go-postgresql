package handler

import (
	"encoding/json"
	"fmt"
	"github.com/carlosd-ss/go-postgresql/go-postgres/net-http/models"
	"github.com/carlosd-ss/go-postgresql/go-postgres/net-http/repo"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		repo.InsertUser(models.User{
			ID:       0,
			Name:     "",
			Location: "",
			Age:      0,
		})
	}

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("tokenjwt", "eyJhbGciOi-RkOM8Hjc5DYNJuqyEy3gvy_IMjcu2w-hl2yHilvPNP_UK0ocUxaKdsD5oS5fV-TYlfH_k")
	w.Header().Set("tokenmsg", "use o token para acessar os endpoints!")
	w.WriteHeader(201)

}

// GetUser will return a single user by its id
func GetUser(w http.ResponseWriter, r *http.Request) {

	// GetUser will return a single user by its id
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// get the userid from the request params, key is "id"
	log.Println(r.URL.Path)
	partes := strings.Split(r.URL.Path, "/")

	var id, err = strconv.Atoi(partes[3])

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return

	}

	// call the getUser function with user id to retrieve a single user
	user, err := repo.GetUser(int64(id))
	if err != nil {

		log.Fatalf("Unable to get user. %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return

	}

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(user)
}

// GetAllUser will return all the users
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "GET" {
		repo.GetAllUsers()

	}

	users, err := repo.GetAllUsers()

	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}

	// send all the users as response
	json.NewEncoder(w).Encode(users)
}

// UpdateUser update user's detail in the postgres db
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// get the userid from the request params, key is "id"
	log.Println(r.URL.Path)
	partes := strings.Split(r.URL.Path, "/")
	var id, err = strconv.Atoi(partes[3])

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return

	}
	// call the getUser function with user id to retrieve a single user
	user, err := repo.GetUser(int64(id))
	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// send the response
	json.NewEncoder(w).Encode(user)

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

}

// DeleteUser delete user's detail in the postgres db
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// get the userid from the request params, key is "id"
	log.Println(r.URL.Path)
	partes := strings.Split(r.URL.Path, "/")
	var id, err = strconv.Atoi(partes[3])

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return

	}
	// call the getUser function with user id to retrieve a single user
	user, err := repo.DeleteUser(int64(id))
	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// send the response
	json.NewEncoder(w).Encode(user)

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

}

// Router is exported and used in main.go
func Router() {

	http.HandleFunc("/api/user/", GetUser)
	http.HandleFunc("/api/user", GetAllUser)
	http.HandleFunc("/api/user", CreateUser)
	http.HandleFunc("/api/user/", UpdateUser)
	http.HandleFunc("/api/user/", DeleteUser)

}

func Server() {
	Router()

	fmt.Println("Server rodando na porta 8080...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
