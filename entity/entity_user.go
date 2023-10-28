package entity

import (
	"context"
	"github.com/gin-gonic/gin"
)

type User struct {
	Base
	NickName    string `db:"nick_name"`
	Email       string `db:"email"`
	Provider    string `db:"provider"`
	FirebaseUID string `db:"firebase_uid"`
	Code        string `db:"code"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	FindByID(ctx context.Context, id BinaryUUID) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	DeleteUser(ctx context.Context, id BinaryUUID) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByCode(ctx context.Context, friendCode string) (*User, error)
}

type UserService interface {
	ReadUser(ctx context.Context, req ReadUserRequest) (*ReadUserResponse, error)
	UpdateUser(ctx context.Context, req UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req DeleteUserRequest) error
	OAuthLoginUser(ctx context.Context, req OAuthLoginUserRequest) (*OAuthLoginUserResponse, error)
}

type UserController interface {
	ReadUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	OAuthLoginUser(c *gin.Context)
}
