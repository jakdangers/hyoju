package entity

import (
	"context"
	"cryptoChallenges/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	Base
	Name     string `db:"name"`
	Email    string `db:"email"`
	UserID   string `db:"user_id"`
	Password string `db:"password"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	ReadUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type UserService interface {
	CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.CreateUserResponse, error)
	ReadUser(ctx context.Context, req dto.ReadUserRequest) (*dto.ReadUserResponse, error)
	UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (*dto.UpdateUserResponse, error)
}

type UserController interface {
	CreateUser(c *gin.Context)
	ReadUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}
