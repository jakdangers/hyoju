package entity

import (
	"context"
	"github.com/gin-gonic/gin"
	"pixelix/dto"
)

type User struct {
	Base
	NickName    string `db:"nick_name"`
	Email       string `db:"email"`
	Provider    string `db:"provider"`
	FirebaseUID string `db:"firebase_uid"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	FindByID(ctx context.Context, id BinaryUUID) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	DeleteUser(ctx context.Context, id BinaryUUID) error
	FindByEmail(ctx context.Context, email string) (*User, error)
}

type UserService interface {
	ReadUser(ctx context.Context, req dto.ReadUserRequest) (*dto.ReadUserResponse, error)
	UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (*dto.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req dto.DeleteUserRequest) error
	OAuthLoginUser(ctx context.Context, req dto.OAuthLoginUserRequest) (*dto.OAuthLoginUserResponse, error)
}

type UserController interface {
	ReadUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	OAuthLoginUser(c *gin.Context)
}
