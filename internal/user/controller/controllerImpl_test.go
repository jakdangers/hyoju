package controller

import (
	"cryptoChallenges/mocks"
	"cryptoChallenges/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type userControllerTestSuite struct {
	suite.Suite
	router         *gin.Engine
	log            log.Logger
	userService    *mocks.UserService
	userController *mocks.UserController
}

func (us *userControllerTestSuite) SetupTest() {
	us.router = gin.Default()
	us.log = mocks.NewLogger(us.T())
	us.userService = mocks.NewUserService(us.T())
	us.userController = mocks.NewUserController(us.T())
	Routes(us.router, us.userController)
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(userControllerTestSuite))
}

func (us *userControllerTestSuite) Test_userController_CreateUser() {
	tests := []struct {
		name  string
		mock  func()
		given func()
		want  func()
	}{
		{
			name: "성공-기본",
			mock: func() {
				us.userService.EXPECT().CreateUser(mock.Anything).Return("create", nil)
			},
			given: nil,
			want:  nil,
		},
	}
	for _, tt := range tests {
		us.T().Run(tt.name, func(t *testing.T) {
			tt.mock()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/users", nil)
			us.router.ServeHTTP(rec, req)
			us.Equal(http.StatusOK, rec.Code)
			us.Equal("create", rec.Body.String())
		})
	}
}
