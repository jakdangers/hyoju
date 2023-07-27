package service

import (
	"context"
	"cryptoChallenges/internal/user/repository"
	"github.com/google/uuid"
)

type userService struct {
	repo repository.UserRepository
}

func New(repo repository.UserRepository) *userService {
	return &userService{repo: repo}
}

var _ UserService = (*userService)(nil)

func (us *userService) CreateUser(ctx context.Context) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (us *userService) ReadUser(ctx context.Context) (string, error) {
	us.repo.ReadUser(ctx, uuid.UUID{})
	return "", nil
}
