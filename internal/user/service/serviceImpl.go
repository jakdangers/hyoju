package service

import "cryptoChallenges/internal/user/repository"

type userService struct {
	repo repository.UserRepository
}

func New(repo repository.UserRepository) *userService {
	return &userService{repo: repo}
}

func (us *userService) GetUsers() (string, error) {
	return us.repo.GetUsers()
}
