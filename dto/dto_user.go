package dto

import (
	"cryptoChallenges/pkg/errors"
	"github.com/google/uuid"
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
	const op errors.Op = "/user/controller/valid"

	if ur.Name == "" || ur.Email == "" || ur.Password == "" {
		return errors.E(op, errors.Invalid, "invalid input")
	}

	return nil
}

type UpdateUserResponse struct {
	ID     string `json:"ID"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	UserID string `json:"userID"`
}
