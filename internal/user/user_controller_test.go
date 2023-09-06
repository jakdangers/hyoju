package user

import (
	"bytes"
	"cryptoChallenges/dto"
	"cryptoChallenges/entity"
	"cryptoChallenges/entity/mocks"
	"cryptoChallenges/pkg/errors"
	"cryptoChallenges/pkg/log"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type userControllerTestSuite struct {
	suite.Suite
	router         *gin.Engine
	log            log.Logger
	userService    *mocks.UserService
	userController entity.UserController
}

func (us *userControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	us.router = gin.Default()
	us.log = mocks.NewLogger(us.T())
	us.userService = mocks.NewUserService(us.T())
	us.userController = NewUserController(us.userService, us.log)
	Routes(us.router, us.userController)
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(userControllerTestSuite))
}

func (us *userControllerTestSuite) Test_userController_CreateUser() {
	tests := []struct {
		name   string
		input  func() *bytes.Reader
		mock   func()
		status int
	}{
		{
			name: "성공-기본",
			input: func() *bytes.Reader {
				user := dto.CreateUserRequest{
					Name:     "cryptoChallenge",
					Email:    "cryptoChallenge@gmail.com",
					UserID:   "cryptoChallenge",
					Password: "cryptoChallenge",
				}
				jsonData, _ := json.Marshal(user)
				return bytes.NewReader(jsonData)
			},
			mock: func() {
				us.userService.On("CreateUser", mock.Anything, dto.CreateUserRequest{
					Name:     "cryptoChallenge",
					Email:    "cryptoChallenge@gmail.com",
					UserID:   "cryptoChallenge",
					Password: "cryptoChallenge",
				}).Return(dto.CreateUserResponse{}, nil).Once()
			},
			status: http.StatusOK,
		},
		{
			name: "실패-기본",
			input: func() *bytes.Reader {
				user := dto.CreateUserRequest{
					Name:     "cryptoChallenge",
					Email:    "cryptoChallenge@gmail.com",
					UserID:   "cryptoChallenge",
					Password: "cryptoChallenge",
				}
				jsonData, _ := json.Marshal(user)
				return bytes.NewReader(jsonData)
			},
			mock: func() {
				us.userService.EXPECT().CreateUser(mock.Anything, dto.CreateUserRequest{
					Name:     "cryptoChallenge",
					Email:    "cryptoChallenge@gmail.com",
					UserID:   "cryptoChallenge",
					Password: "cryptoChallenge",
				}).Return(&dto.CreateUserResponse{}, &errors.Error{Kind: errors.Internal}).Once()
			},
			status: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		us.T().Run(tt.name, func(t *testing.T) {
			// input
			tt.mock()
			req, _ := http.NewRequest(http.MethodPost, "/users", tt.input())
			req.Header.Set("Content-Type", "application/json")

			// when
			rec := httptest.NewRecorder()
			us.router.ServeHTTP(rec, req)

			// then
			us.Equal(tt.status, rec.Code)
			us.userService.AssertExpectations(t)
		})
	}
}

func (us *userControllerTestSuite) Test_userController_GetUser() {
	testUserID := "998c084a-9982-4c24-9663-4f24e2e3db36"

	tests := []struct {
		name   string
		input  func() string
		mock   func()
		status int
	}{
		{
			name: "pass ID로 조회",
			input: func() string {
				params := url.Values{}
				params.Add("ID", testUserID)
				return params.Encode()
			},
			mock: func() {
				var res dto.ReadUserResponse
				err := faker.FakeData(&res)
				us.NoError(err)
				us.userService.EXPECT().ReadUser(mock.Anything, dto.ReadUserRequest{
					ID: testUserID,
				}).Return(&res, nil)
			},
			status: http.StatusOK,
		},
	}
	for _, tt := range tests {
		us.T().Run(tt.name, func(t *testing.T) {
			// input
			tt.mock()
			req, _ := http.NewRequest(http.MethodGet, "/users", nil)
			req.URL.RawQuery = tt.input()

			// when
			rec := httptest.NewRecorder()
			us.router.ServeHTTP(rec, req)

			// then
			us.Equal(tt.status, rec.Code)
			us.userService.AssertExpectations(t)
		})
	}
}

func (us *userControllerTestSuite) Test_userController_UpdateUser() {
	testUserID := uuid.New()

	tests := []struct {
		name   string
		input  func() *bytes.Reader
		mock   func()
		status int
	}{
		{
			name: "PASS 기본",
			input: func() *bytes.Reader {
				req := dto.UpdateUserRequest{
					ID:       testUserID.String(),
					Name:     "cryptoChallenge",
					Email:    "cryptoChallenge@gmail.com",
					Password: "cryptoChallenge",
				}
				jsonData, _ := json.Marshal(req)
				return bytes.NewReader(jsonData)
			},
			mock: func() {
				us.userService.EXPECT().UpdateUser(mock.Anything, dto.UpdateUserRequest{
					ID:       testUserID.String(),
					Name:     "cryptoChallenge",
					Email:    "cryptoChallenge@gmail.com",
					Password: "cryptoChallenge",
				}).Return(&dto.UpdateUserResponse{
					ID:     testUserID.String(),
					UserID: "cryptoChallenge",
					Name:   "cryptoChallenge",
					Email:  "cryptoChallenge@gmail.com",
				}, nil).Once()
			},
			status: http.StatusOK,
		},
		{
			name: "FAIL 기본",
			input: func() *bytes.Reader {
				req := dto.UpdateUserRequest{
					ID:       testUserID.String(),
					Name:     "",
					Email:    "",
					Password: "",
				}
				jsonData, _ := json.Marshal(req)
				return bytes.NewReader(jsonData)
			},
			mock: func() {
			},
			status: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		us.T().Run(tt.name, func(t *testing.T) {
			// input
			tt.mock()
			req, _ := http.NewRequest(http.MethodPut, "/users", tt.input())
			req.Header.Set("Content-Type", "application/json")

			// when
			rec := httptest.NewRecorder()
			us.router.ServeHTTP(rec, req)

			// then
			us.Equal(tt.status, rec.Code)
			us.userService.AssertExpectations(t)
		})
	}
}
