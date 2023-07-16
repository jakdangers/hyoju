package service

import "context"

type UserService interface {
	ReadUser(ctx context.Context) (string, error)
}
