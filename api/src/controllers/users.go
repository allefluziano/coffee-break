package controllers

import "net/http"

//CreateUser insert a user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("creating user!"))
}

//GetUsers list all users in database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("searching all users!"))
}

//GetUser show a user informations
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("searching a user!"))
}

//UpdateUser update a user informations
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updating a user!"))
}

//DeleteUser delete a user in database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleting a user!"))
}
