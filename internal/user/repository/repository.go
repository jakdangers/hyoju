package repository

import (
	"context"
	"cryptoChallenges/internal/user/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	ReadUser(ctx context.Context, id uuid.UUID) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
