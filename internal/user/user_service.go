package user

import (
	"context"
	"github.com/google/uuid"
	"pixelix/dto"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
)

type userService struct {
	repository entity.UserRepository
}

func NewUserService(repo entity.UserRepository) *userService {
	return &userService{repository: repo}
}

var _ entity.UserService = (*userService)(nil)

func (us *userService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	const op cerrors.Op = "user/service/createUser"
	findUser, err := us.repository.ReadUser(ctx, &entity.User{UserID: req.UserID})
	if err != nil {
		return nil, err
	}
	if findUser != nil {
		return nil, cerrors.E(op, cerrors.Exist, "userID already exists")
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
	const op cerrors.Op = "user/service/readUser"

	userID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
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
		return nil, cerrors.E(op, cerrors.Internal, err)
	}
	if user == nil {
		return nil, cerrors.E(op, cerrors.NotExist)
	}

	return &dto.ReadUserResponse{
		ID:     user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
		UserID: user.UserID,
	}, nil
}

func (us *userService) UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (*dto.UpdateUserResponse, error) {
	const op cerrors.Op = "user/service/updateUser"

	userID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	findUser, err := us.repository.ReadUser(ctx, &entity.User{Base: entity.Base{
		ID: userID,
	}})
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}
	if findUser == nil {
		return nil, cerrors.E(op, cerrors.NotExist)
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
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	return &dto.UpdateUserResponse{
		ID:     user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
		UserID: user.UserID,
	}, nil
}

func (us *userService) DeleteUser(ctx context.Context, req dto.DeleteUserRequest) error {
	//TODO implement me
	panic("implement me")
}
