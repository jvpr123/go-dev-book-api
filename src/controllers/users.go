package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Route to create user"))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Route to list all users"))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Route to user by ID"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Route to update user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Route to delete user"))
}
