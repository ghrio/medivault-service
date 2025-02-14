package handler

import (
	"medivault-service/internal/services"
	"net/http"
)

type UserHandler struct {
	Service services.UserService
}

// CreateUser creates a new user
// @Summary      Create a new user
// @Description  Register a new user in the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User data"
// @Success      201   {object}  models.User
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /users [post]
func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

}

// GetUserByID retrieves a user by ID
// @Summary      Get user by ID
// @Description  Retrieve a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "User ID"
// @Success      200   {object}  models.User
// @Failure      400   {object}  ErrorResponse
// @Failure      404   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /users/{id} [get]
func (u *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {}
