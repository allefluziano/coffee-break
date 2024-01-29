package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateUser insert a user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUsersRepository(db)
	userID, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Inserted Id: %d", userID)))
}

// GetUsers list all users in database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("searching all users!"))
}

// GetUser show a user informations
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("searching a user!"))
}

// UpdateUser update a user informations
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updating a user!"))
}

// DeleteUser delete a user in database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleting a user!"))
}
