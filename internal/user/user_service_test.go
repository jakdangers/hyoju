package user

import (
	"context"
	"cryptoChallenges/dto"
	"cryptoChallenges/entity"
	"cryptoChallenges/entity/mocks"
	"cryptoChallenges/pkg/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type userServiceTestSuite struct {
	suite.Suite
	repository *mocks.UserRepository
	service    entity.UserService
}

func (us *userServiceTestSuite) SetupTest() {
	us.repository = mocks.NewUserRepository(us.T())
	us.service = NewUserService(us.repository)
}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(userServiceTestSuite))
}

func (us *userServiceTestSuite) Test_userService_CreateUser() {
	testUserID := uuid.New()
	tests := []struct {
		name    string
		input   dto.CreateUserRequest
		ctx     context.Context
		mock    func()
		want    *dto.CreateUserResponse
		wantErr bool
	}{
		{
			name: "pass 기본 생성",
			input: dto.CreateUserRequest{
				Name:     "cryptoChallenges",
				Email:    "cryptochallenges@gmail.com",
				UserID:   "cryptoChallenges",
				Password: "password",
			},
			ctx: context.Background(),
			mock: func() {
				user := entity.User{
					Name:     "cryptoChallenges",
					Email:    "cryptochallenges@gmail.com",
					UserID:   "cryptoChallenges",
					Password: "password",
				}
				us.repository.EXPECT().ReadUser(mock.Anything, mock.Anything).Return(nil, nil).Once()
				us.repository.EXPECT().CreateUser(mock.Anything, &user).Return(&entity.User{
					Base: entity.Base{
						ID: testUserID,
					},
					Name:     "cryptoChallenges",
					Email:    "cryptochallenges@gmail.com",
					UserID:   "cryptoChallenges",
					Password: "password",
				}, nil).Once()
			},
			want: &dto.CreateUserResponse{
				ID:     testUserID.String(),
				Name:   "cryptoChallenges",
				Email:  "cryptochallenges@gmail.com",
				UserID: "cryptoChallenges",
			},
			wantErr: false,
		},
		{
			name: "fail repository error",
			input: dto.CreateUserRequest{
				Name:     "cryptoChallenges",
				Email:    "cryptoChallenges@gmail.com",
				UserID:   "cryptoChallenges",
				Password: "cryptoChallenges",
			},
			ctx: context.Background(),
			mock: func() {
				us.repository.EXPECT().ReadUser(mock.Anything, &entity.User{
					UserID: "cryptoChallenges",
				}).Return(nil, errors.E(errors.Internal)).Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fail UserID 중복",
			input: dto.CreateUserRequest{
				Name:     "cryptoChallenges",
				Email:    "cryptoChallenges@gmail.com",
				UserID:   "cryptoChallenges",
				Password: "cryptoChallenges",
			},
			ctx: context.Background(),
			mock: func() {
				us.repository.EXPECT().ReadUser(mock.Anything, &entity.User{
					UserID: "cryptoChallenges",
				}).Return(&entity.User{
					Name:     "cryptoChallenges",
					Email:    "cryptoChallenges@gmail.com",
					UserID:   "cryptoChallenges",
					Password: "cryptoChallenges",
				}, nil).Once()
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		us.T().Run(tt.name, func(t *testing.T) {
			tt.mock()
			result, err := us.service.CreateUser(tt.ctx, tt.input)
			us.Equal(tt.want, result)
			us.Equal(tt.wantErr, err != nil)
		})
	}
}

func (us *userServiceTestSuite) Test_userService_ReadUser() {
	testUserID := uuid.New()

	tests := []struct {
		name    string
		input   dto.ReadUserRequest
		ctx     context.Context
		mock    func()
		want    *dto.ReadUserResponse
		wantErr bool
	}{
		{
			name: "pass - ID로 조회",
			input: dto.ReadUserRequest{
				ID: testUserID.String(),
			},
			ctx: context.Background(),
			mock: func() {
				us.repository.EXPECT().ReadUser(mock.Anything, &entity.User{
					Base: entity.Base{
						ID: testUserID,
					},
				}).Return(&entity.User{
					Base: entity.Base{
						ID: testUserID,
					},
					Name:     "cryptoChallenges",
					Email:    "cryptoChallenges@gmail.com",
					UserID:   "cryptoChallenges",
					Password: "",
				}, nil)
			},
			want: &dto.ReadUserResponse{
				ID:     testUserID.String(),
				Name:   "cryptoChallenges",
				Email:  "cryptoChallenges@gmail.com",
				UserID: "cryptoChallenges",
			},
		},
	}
	for _, tt := range tests {
		us.T().Run(tt.name, func(t *testing.T) {
			tt.mock()
			res, err := us.service.ReadUser(tt.ctx, tt.input)
			us.Equal(tt.want, res)
			us.Equal(tt.wantErr, err != nil)
		})
	}
}

func (us *userServiceTestSuite) Test_userService_UpdateUser() {
	testUserID := uuid.New()

	tests := []struct {
		name    string
		input   dto.UpdateUserRequest
		ctx     context.Context
		mock    func()
		want    *dto.UpdateUserResponse
		wantErr bool
	}{
		{
			name: "PASS 기본",
			input: dto.UpdateUserRequest{
				ID:       "",
				Name:     "",
				Email:    "",
				Password: "",
			},
			ctx: context.Background(),
			mock: func() {
				us.repository.EXPECT().UpdateUser(mock.Anything, &entity.User{
					Base: entity.Base{
						ID: testUserID,
					},
					Name:     "modified_cryptoChallenges",
					Email:    "modified_cryptoChallenges@gmail.com",
					Password: "modified_cryptoChallenges",
				}).Return(&entity.User{
					Base:     entity.Base{},
					Name:     "modified_cryptoChallenges",
					Email:    "modified_cryptoChallenges@gmail.com",
					UserID:   "cryptoChallenges",
					Password: "modified_cryptoChallenges",
				}, nil).Once()
			},
			want: &dto.UpdateUserResponse{
				ID:     testUserID.String(),
				Name:   "cryptoChallenges",
				Email:  "cryptoChallenges@gmail.com",
				UserID: "cryptoChallenges",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		us.T().Run(tt.name, func(t *testing.T) {
			tt.mock()
			res, err := us.service.UpdateUser(tt.ctx, tt.input)
			us.Equal(tt.want, res)
			us.Equalf(tt.wantErr, err != nil, err.Error())
		})
	}
}

struct {
	struct {
		struct{
			s
}
j
}
}

{
	인간 {
		룩코 {

}
}
}