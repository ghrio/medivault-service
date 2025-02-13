package handler

import (
	"medivault-service/internal/services"
	"net/http"
)

type UserHandler struct {
	Service services.UserService
}

//Create User

func CreateUser(w http.ResponseWriter, r *http.Request) {}

// Get UserByID
func GetUserByID(w http.ResponseWriter, r *http.Request) {}

// Get All Users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {}

// Update User
func UpdateUser(w http.ResponseWriter, r *http.Request) {}

// Delete User
func DeleteUser(w http.ResponseWriter, r *http.Request) {}
