package repository

import (
	"cryptoChallenges/internal/user/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user *entity.User) (entity.User, error)
	ReadUser(id uuid.UUID) (entity.User, error)
	UpdateUser(user *entity.User) (entity.User, error)
	DeleteUser(id uuid.UUID) error
}
