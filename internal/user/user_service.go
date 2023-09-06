package user

import (
	"context"
	"cryptoChallenges/dto"
	"cryptoChallenges/entity"
	"cryptoChallenges/pkg/errors"
	"github.com/google/uuid"
)

type userService struct {
	repository entity.UserRepository
}

func NewUserService(repo entity.UserRepository) *userService {
	return &userService{repository: repo}
}

var _ entity.UserService = (*userService)(nil)

func (us *userService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	const op errors.Op = "user/service/createUser"
	findUser, err := us.repository.ReadUser(ctx, &entity.User{UserID: req.UserID})
	if err != nil {
		return nil, err
	}
	if findUser != nil {
		return nil, errors.E(op, errors.Exist, "userID already exists")
	}

	user, err := us.repository.CreateUser(ctx, &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		UserID:   req.UserID,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &dto.CreateUserResponse{
		ID:     user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
		UserID: user.UserID,
	}, nil
}

func (us *userService) ReadUser(ctx context.Context, req dto.ReadUserRequest) (*dto.ReadUserResponse, error) {
	const op errors.Op = "user/service/readUser"

	userID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, errors.E(op, errors.Internal, err)
	}

	user, err := us.repository.ReadUser(ctx, &entity.User{
		Base: entity.Base{
			ID: userID,
		},
		Name:   req.Name,
		Email:  req.Email,
		UserID: req.UserID,
	})
	if err != nil {
		return nil, errors.E(op, errors.Internal, err)
	}

	return &dto.ReadUserResponse{
		ID:     user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
		UserID: user.UserID,
	}, nil
}

func (us *userService) UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (*dto.UpdateUserResponse, error) {
	const op errors.Op = "user/service/updateUser"

	userID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, errors.E(op, errors.Internal, err, "정상요청인데 뭐가 서버에서 사망이야")
	}

	user, err := us.repository.UpdateUser(ctx, &entity.User{
		Base: entity.Base{
			ID: userID,
		},
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, errors.E(op, errors.Internal, err)
	}

	return &dto.UpdateUserResponse{
		ID:     user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
		UserID: user.UserID,
	}, nil
}
