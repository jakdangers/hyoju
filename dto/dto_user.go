package dto

import (
	"github.com/google/uuid"
	"pixelix/pkg/cerrors"
)

type UserDTO struct {
	ID     uuid.UUID `json:"ID"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	UserID string    `json:"userID"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserID   string `json:"userID"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	ID     string `json:"ID"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	UserID string `json:"userID"`
}

type ReadUserRequest struct {
	ID     string `form:"ID"`
	Name   string `form:"name"`
	Email  string `form:"email"`
	UserID string `form:"userID"`
}

type ReadUserResponse struct {
	ID     string `form:"ID"`
	Name   string `form:"name"`
	Email  string `form:"email"`
	UserID string `form:"userID"`
}

type UpdateUserRequest struct {
	ID       string `json:"ID"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ur UpdateUserRequest) Valid() error {
	const op cerrors.Op = "/user/controller/valid"

	if ur.Name == "" || ur.Email == "" || ur.Password == "" {
		return cerrors.E(op, cerrors.Invalid, "invalid input")
	}

	return nil
}

type UpdateUserResponse struct {
	ID     string `json:"ID"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	UserID string `json:"userID"`
}

type DeleteUserRequest struct {
	ID string `json:"ID" uri:"ID"`
}
