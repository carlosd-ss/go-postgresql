package handler

import (
	"encoding/json"
	"fmt"
	"github.com/carlosd-ss/go-postgresql/go-postgres/gorila-mux/models"
	"github.com/carlosd-ss/go-postgresql/go-postgres/gorila-mux/repo"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
	"log"
	"net/http"

	"strconv"
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("tokenjwt", "eyJhbGciOi-RkOM8Hjc5DYNJuqyEy3gvy_IMjcu2w-hl2yHilvPNP_UK0ocUxaKdsD5oS5fV-TYlfH_k")
	w.Header().Set("tokenmsg", "use o token para acessar os endpoints!")
	w.WriteHeader(201)

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	insertID := repo.InsertUser(user)

	res := response{
		ID:      insertID,
		Message: "User created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	user, err := repo.GetUser(int64(id))

	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	json.NewEncoder(w).Encode(user)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	users, err := repo.GetAllUsers()

	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}

	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	var user models.User

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := repo.UpdateUser(int64(id), user)

	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := repo.DeleteUser(int64(id))

	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
