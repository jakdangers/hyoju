package user

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"pixelix/entity"
	"pixelix/mocks"
	"pixelix/pkg/cerrors"
	"pixelix/pkg/logger"
	"testing"
)

type userControllerTestSuite struct {
	router         *gin.Engine
	log            logger.Logger
	userService    *mocks.UserService
	userController entity.UserController
}

func setupUserControllerTestSuite(t *testing.T) userControllerTestSuite {
	var us userControllerTestSuite

	gin.SetMode(gin.TestMode)
	us.router = gin.Default()
	us.userService = mocks.NewUserService(t)
	us.userController = NewUserController(us.userService, us.log)
	RegisterRoutes(us.router, us.userController)

	return us
}

func Test_userController_ReadUser(t *testing.T) {
	testUserID := "998c084a-9982-4c24-9663-4f24e2e3db36"
	us := setupUserControllerTestSuite(t)

	tests := []struct {
		name   string
		input  func() string
		mock   func()
		status int
	}{
		{
			name: "PASS 존재하는 userID 조회",
			input: func() string {
				params := url.Values{}
				params.Add("ID", testUserID)
				return params.Encode()
			},
			mock: func() {
				us.userService.EXPECT().ReadUser(mock.Anything, entity.ReadUserRequest{
					ID: testUserID,
				}).Return(&entity.ReadUserResponse{
					ID: testUserID,
				}, nil).Once()
			},
			status: http.StatusOK,
		},
		{
			name: "PASS 존재하지 않는 userID 조회",
			input: func() string {
				params := url.Values{}
				params.Add("ID", testUserID)
				return params.Encode()
			},
			mock: func() {
				us.userService.EXPECT().ReadUser(mock.Anything, entity.ReadUserRequest{
					ID: testUserID,
				}).Return(nil, cerrors.E(cerrors.Invalid)).Once()
			},
			status: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// input
			tt.mock()
			req, _ := http.NewRequest(http.MethodGet, "/users", nil)
			req.URL.RawQuery = tt.input()

			// when
			rec := httptest.NewRecorder()
			us.router.ServeHTTP(rec, req)

			// then
			assert.Equal(t, tt.status, rec.Code)
			us.userService.AssertExpectations(t)
		})
	}
}

func Test_userController_UpdateUser(t *testing.T) {
	us := setupUserControllerTestSuite(t)
	testUserID := uuid.New()

	tests := []struct {
		name   string
		input  func() *bytes.Reader
		mock   func()
		status int
	}{
		{
			name: "PASS 존재하는 userID 수정",
			input: func() *bytes.Reader {
				req := entity.UpdateUserRequest{
					ID:       testUserID.String(),
					NickName: "modified_nickName",
				}
				jsonData, _ := json.Marshal(req)
				return bytes.NewReader(jsonData)
			},
			mock: func() {
				us.userService.EXPECT().UpdateUser(mock.Anything, entity.UpdateUserRequest{
					ID:       testUserID.String(),
					NickName: "modified_nickName",
				}).Return(&entity.UpdateUserResponse{
					ID:       testUserID.String(),
					Email:    "blipix@blipix.com",
					NickName: "modified_nickName",
				}, nil).Once()
			},
			status: http.StatusOK,
		},
		{
			name: "FAIL 존재하지 않는 userID 수정",
			input: func() *bytes.Reader {
				req := entity.UpdateUserRequest{
					ID:       testUserID.String(),
					NickName: "modified_nickName",
				}
				jsonData, _ := json.Marshal(req)
				return bytes.NewReader(jsonData)
			},
			mock: func() {
				us.userService.EXPECT().UpdateUser(mock.Anything, entity.UpdateUserRequest{
					ID:       testUserID.String(),
					NickName: "modified_nickName",
				}).Return(nil, cerrors.E(cerrors.Invalid)).Once()
			},
			status: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// input
			tt.mock()
			req, _ := http.NewRequest(http.MethodPut, "/users", tt.input())
			req.Header.Set("Content-Type", "application/json")

			// when
			rec := httptest.NewRecorder()
			us.router.ServeHTTP(rec, req)

			// then
			assert.Equal(t, tt.status, rec.Code)
			us.userService.AssertExpectations(t)
		})
	}
}

func Test_userController_DeleteUser(t *testing.T) {
	us := setupUserControllerTestSuite(t)
	testUserID := uuid.New()

	tests := []struct {
		name   string
		input  func() string
		mock   func()
		status int
	}{
		{
			name: "PASS 존재하는 userID",
			input: func() string {
				path, _ := url.JoinPath("/users", testUserID.String())
				return path
			},
			mock: func() {
				us.userService.EXPECT().DeleteUser(mock.Anything, entity.DeleteUserRequest{
					ID: testUserID.String(),
				}).Return(nil).Once()
			},
			status: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			req, _ := http.NewRequest(http.MethodDelete, tt.input(), nil)

			rec := httptest.NewRecorder()
			us.router.ServeHTTP(rec, req)

			assert.Equal(t, tt.status, rec.Code)
			us.userService.AssertExpectations(t)
		})
	}
}

func Test_userController_OAuthLoginUser(t *testing.T) {
	us := setupUserControllerTestSuite(t)

	tests := []struct {
		name  string
		input func() *bytes.Reader
		mock  func()
		want  int
	}{
		{
			name: "PASS OAuth 로그인",
			input: func() *bytes.Reader {
				req := entity.OAuthLoginUserRequest{
					Email:       "blipix@blipix.com",
					FirebaseUID: "firebaseUID",
					Provider:    "blipix",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func() {
				us.userService.EXPECT().OAuthLoginUser(mock.Anything, entity.OAuthLoginUserRequest{
					Email:       "blipix@blipix.com",
					FirebaseUID: "firebaseUID",
					Provider:    "blipix",
				}).Return(&entity.OAuthLoginUserResponse{
					AccessToken: "accessToken",
				}, nil)
			},
			want: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			req, _ := http.NewRequest(http.MethodPost, "/users/login", tt.input())
			req.Header.Set("Content-Type", "application/json")

			// when
			rec := httptest.NewRecorder()
			us.router.ServeHTTP(rec, req)

			// then
			assert.Equal(t, tt.want, rec.Code)
			us.userService.AssertExpectations(t)
		})
	}
}
