package controller

import (
	"cryptoChallenges/internal/user/service"
	"cryptoChallenges/mocks"
	"cryptoChallenges/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"testing"
)

type userControllerTestSuite struct {
	suite.Suite
	router         *gin.Engine
	log            log.Logger
	userService    service.UserService
	userController UserController
}

func (us *userControllerTestSuite) SetupTest() {
	us.router = gin.Default()
	us.log = mocks.NewLogger(us.T())
	us.userService = mocks.NewUserService(us.T())
	us.userController = New(us.userService, us.log)
}

//func (us *userControllerTestSuite) Test_userController_CreateUser() {
//	tests := []struct {
//		name string
//	}{
//		{
//			name: "success case",
//		},
//	}
//	for _, tt := range tests {
//		w := httptest.NewRecorder()
//		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
//		us.router.ServeHTTP(w, req)
//		assert.Equal(t, 200, w.Code)
//		assert.Equal(t, "pong", w.Body.String())
//	}
//}

func Test_userController_GetUser(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {})
	}
}
