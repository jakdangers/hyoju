package service

import "context"

type UserService interface {
	CreateUser(ctx context.Context, createUserRequest)
	ReadUser(ctx context.Context) (string, error)
}
