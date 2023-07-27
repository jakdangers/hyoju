package service

import "context"

type UserService interface {
	CreateUser(ctx context.Context) (string, error)
	ReadUser(ctx context.Context) (string, error)
}
