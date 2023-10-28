package user

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
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

func (us *userService) ReadUser(ctx context.Context, req entity.ReadUserRequest) (*entity.ReadUserResponse, error) {
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

	return &entity.ReadUserResponse{
		ID:       user.ID.String(),
		Email:    user.Email,
		NickName: user.NickName,
	}, nil
}

func (us *userService) UpdateUser(ctx context.Context, req entity.UpdateUserRequest) (*entity.UpdateUserResponse, error) {
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

	return &entity.UpdateUserResponse{
		ID:       user.ID.String(),
		Email:    user.Email,
		NickName: user.NickName,
	}, nil
}

func (us *userService) DeleteUser(ctx context.Context, req entity.DeleteUserRequest) error {
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

func (us *userService) OAuthLoginUser(ctx context.Context, req entity.OAuthLoginUserRequest) (*entity.OAuthLoginUserResponse, error) {
	const op cerrors.Op = "user/service/OAuthLoginUser"

	user, err := us.repository.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	// 유저가 존재하지 않는 경우 생성
	if user == nil {
		userID := entity.BinaryUUIDNew()
		for {
			findUser, err := us.repository.FindByCode(ctx, userID.ToCode())
			if err != nil {
				return nil, cerrors.E(op, cerrors.Internal, err)
			}
			if findUser == nil {
				break
			}
			userID = entity.BinaryUUIDNew()
		}
		user, err = us.repository.CreateUser(ctx, &entity.User{
			Base: entity.Base{
				ID: userID,
			},
			NickName:    req.Email,
			Email:       req.Email,
			Provider:    req.Provider,
			FirebaseUID: req.FirebaseUID,
			Code:        userID.ToCode(),
		})
		if err != nil {
			return nil, cerrors.E(op, cerrors.Internal, err)
		}
	}

	return &entity.OAuthLoginUserResponse{
		ID:          user.ID.String(),
		NickName:    user.NickName,
		Email:       user.Email,
		FriendCode:  user.Code,
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
