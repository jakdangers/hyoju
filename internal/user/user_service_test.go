package user

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"pixelix/entity"
	"pixelix/mocks"
	"testing"
)

type serviceTestSuite struct {
	repository *mocks.UserRepository
	service    entity.UserService
}

func setupUserServiceTestSuite(t *testing.T) serviceTestSuite {
	var us serviceTestSuite

	us.repository = mocks.NewUserRepository(t)
	us.service = NewUserService(us.repository)

	return us
}

func Test_userService_ReadUser(t *testing.T) {
	type args struct {
		ctx context.Context
		req entity.ReadUserRequest
	}

	us := setupUserServiceTestSuite(t)
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *entity.ReadUserResponse
		wantErr bool
	}{
		{
			name: "PASS 존재하는 유저 조회",
			args: args{
				ctx: context.Background(),
				req: entity.ReadUserRequest{
					ID: testUserID.String(),
				},
			},
			mock: func() {
				us.repository.EXPECT().FindByID(mock.Anything, testUserID).Return(&entity.User{
					Base: entity.Base{
						ID: testUserID,
					},
					Email:    "blipix@blipix.com",
					NickName: "blipix",
				}, nil).Once()
			},
			want: &entity.ReadUserResponse{
				ID:       testUserID.String(),
				Email:    "blipix@blipix.com",
				NickName: "blipix",
			},
			wantErr: false,
		},
		{
			name: "PASS 존재하지 않는 유저 조회",
			args: args{
				ctx: context.Background(),
				req: entity.ReadUserRequest{
					ID: testUserID.String(),
				},
			},
			mock: func() {
				us.repository.EXPECT().FindByID(mock.Anything, testUserID).Return(nil, nil).Once()
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := us.service.ReadUser(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
			us.repository.AssertExpectations(t)
		})
	}
}

func Test_userService_UpdateUser(t *testing.T) {
	type args struct {
		ctx context.Context
		req entity.UpdateUserRequest
	}

	us := setupUserServiceTestSuite(t)
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *entity.UpdateUserResponse
		wantErr bool
	}{
		{
			name: "PASS 존재하는 userID 수정",
			args: args{
				ctx: context.Background(),
				req: entity.UpdateUserRequest{
					ID:       testUserID.String(),
					NickName: "modified_nickName",
				},
			},
			mock: func() {
				us.repository.EXPECT().FindByID(mock.Anything, testUserID).Return(&entity.User{
					Base: entity.Base{
						ID: testUserID,
					},
					NickName:    "original_nickName",
					Email:       "blipix@blipix.com",
					Provider:    "blipix",
					FirebaseUID: "firebaseUID",
				}, nil).Once()
				us.repository.EXPECT().UpdateUser(mock.Anything, &entity.User{
					Base: entity.Base{
						ID: testUserID,
					},
					NickName:    "modified_nickName",
					Email:       "blipix@blipix.com",
					Provider:    "blipix",
					FirebaseUID: "firebaseUID",
				}).Return(&entity.User{
					Base: entity.Base{
						ID: testUserID,
					},
					NickName:    "modified_nickName",
					Email:       "blipix@blipix.com",
					Provider:    "blipix",
					FirebaseUID: "firebaseUID",
				}, nil).Once()
			},
			want: &entity.UpdateUserResponse{
				ID:       testUserID.String(),
				NickName: "modified_nickName",
				Email:    "blipix@blipix.com",
			},
			wantErr: false,
		},
		{
			name: "FAIL 존재하지 않는 userID 수정",
			args: args{
				ctx: context.Background(),
				req: entity.UpdateUserRequest{
					ID:       testUserID.String(),
					NickName: "modified_nickName",
				},
			},
			mock: func() {
				us.repository.EXPECT().FindByID(mock.Anything, testUserID).Return(nil, nil).Once()
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := us.service.UpdateUser(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
			us.repository.AssertExpectations(t)
		})
	}
}

func Test_userService_OAuthLoginUser(t *testing.T) {
	type args struct {
		ctx context.Context
		req entity.OAuthLoginUserRequest
	}

	us := setupUserServiceTestSuite(t)

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *entity.OAuthLoginUserResponse
		wantErr bool
	}{
		{
			name: "PASS 신규 유저 생성 후 토큰 발생",
			args: args{
				ctx: context.Background(),
				req: entity.OAuthLoginUserRequest{
					Email:       "blipix@blipix.com",
					FirebaseUID: "firebaseUID",
					Provider:    "blipix",
				},
			},
			mock: func() {
				us.repository.EXPECT().FindByEmail(mock.Anything, "blipix@blipix.com").Return(nil, nil).Once()
				us.repository.EXPECT().FindByFriendCode(mock.Anything, mock.Anything).Return(nil, nil).Once()
				us.repository.EXPECT().CreateUser(mock.Anything, mock.Anything).Return(&entity.User{
					NickName:    "blipix@blipix.com",
					Email:       "blipix@blipix.com",
					Provider:    "blipix",
					FirebaseUID: "firebaseUID",
				}, nil).Once()
			},
			want: &entity.OAuthLoginUserResponse{
				NickName:    "blipix@blipix.com",
				Email:       "blipix@blipix.com",
				AccessToken: "test_accessToken",
			},
			wantErr: false,
		},
		{
			name: "PASS 이미 존재하는 유저의 토큰 발생",
			args: args{
				ctx: context.Background(),
				req: entity.OAuthLoginUserRequest{
					Email:       "blipix@blipix.com",
					FirebaseUID: "firebaseUID",
					Provider:    "blipix",
				},
			},
			mock: func() {
				us.repository.EXPECT().FindByEmail(mock.Anything, "blipix@blipix.com").Return(&entity.User{
					NickName: "blipix@blipix.com",
					Email:    "blipix@blipix.com",
				}, nil).Once()
			},
			want: &entity.OAuthLoginUserResponse{
				NickName:    "blipix@blipix.com",
				Email:       "blipix@blipix.com",
				AccessToken: "test_accessToken",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := us.service.OAuthLoginUser(tt.args.ctx, tt.args.req)
			assert.Equal(t, true, cmp.Equal(tt.want, got, cmpopts.IgnoreFields(entity.OAuthLoginUserResponse{}, "MissionID", "FriendCode", "AccessToken")))
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_userService_DeleteUser(t *testing.T) {
	type args struct {
		ctx context.Context
		req entity.DeleteUserRequest
	}

	us := setupUserServiceTestSuite(t)
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "PASS 삭제",
			args: args{
				ctx: context.Background(),
				req: entity.DeleteUserRequest{
					ID: testUserID.String(),
				},
			},
			mock: func() {
				us.repository.EXPECT().DeleteUser(mock.Anything, testUserID).Return(nil).Once()
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := us.service.DeleteUser(tt.args.ctx, tt.args.req)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}
