package user

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
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

func (us *userService) ReadUser(ctx context.Context, req dto.ReadUserRequest) (*dto.ReadUserResponse, error) {
	const op cerrors.Op = "user/service/readUser"

	userID, err := entity.ParseUUID(req.ID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Invalid, err)
	}

	user, err := us.repository.FindByID(ctx, userID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}
	if user == nil {
		return nil, cerrors.E(op, cerrors.Invalid, "user not exist")
	}

	return &dto.ReadUserResponse{
		ID:       user.ID.String(),
		Email:    user.Email,
		NickName: user.NickName,
	}, nil
}

func (us *userService) UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (*dto.UpdateUserResponse, error) {
	const op cerrors.Op = "user/service/updateUser"

	userID, err := entity.ParseUUID(req.ID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Invalid, err)
	}

	findUser, err := us.repository.FindByID(ctx, userID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}
	if findUser == nil {
		return nil, cerrors.E(op, cerrors.Invalid, "user not exist")
	}

	findUser.NickName = req.NickName

	user, err := us.repository.UpdateUser(ctx, findUser)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	return &dto.UpdateUserResponse{
		ID:       user.ID.String(),
		Email:    user.Email,
		NickName: user.NickName,
	}, nil
}

func (us *userService) DeleteUser(ctx context.Context, req dto.DeleteUserRequest) error {
	const op cerrors.Op = "user/service/deleteUser"

	userID, err := entity.ParseUUID(req.ID)
	if err != nil {
		return cerrors.E(op, cerrors.Invalid, err)
	}

	if err := us.repository.DeleteUser(ctx, userID); err != nil {
		return cerrors.E(op, cerrors.Internal, err)
	}

	return nil
}

func (us *userService) OAuthLoginUser(ctx context.Context, req dto.OAuthLoginUserRequest) (*dto.OAuthLoginUserResponse, error) {
	const op cerrors.Op = "user/service/OAuthLoginUser"

	user, err := us.repository.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	// 유저가 존재하지 않는 경우 생성
	if user == nil {
		user, err = us.repository.CreateUser(ctx, &entity.User{
			Base: entity.Base{
				ID: entity.BinaryUUIDNew(),
			},
			Email:       req.Email,
			NickName:    req.Email,
			Provider:    req.Provider,
			FirebaseUID: req.FirebaseUID,
		})
		if err != nil {
			return nil, cerrors.E(op, cerrors.Internal, err)
		}
	}

	return &dto.OAuthLoginUserResponse{
		AccessToken: generateAccessToken(user),
	}, nil
}

func generateAccessToken(user *entity.User) string {
	t := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"iss": "blipix",
			"sub": user.ID,
		})
	accessToken, _ := t.SigningString()

	return accessToken
}
