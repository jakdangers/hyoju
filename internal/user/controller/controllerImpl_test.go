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
				us.userService.EXPECT().CreateUser(mock.Anything).Return("test", nil)
			},
			given: nil,
			want:  nil,
		},
	}
	for _, tt := range tests {
		us.T().Run(tt.name, func(t *testing.T) {
			tt.mock()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/ping", nil)
			us.router.ServeHTTP(w, req)
			//got, err := ur.CreateUser(us.ctx, tt.want)
			//if err == nil {
			//	us.Equal(true, cmp.Equal(tt.want, got, cmpopts.IgnoreFields(entity.User{}, "CreatedAt", "UpdatedAt", "DeletedAt")))
			//}
			//if err != nil {
			//	us.EqualError(err, "user/createUser: internal error: duplicated key not allowed")
			//}
		})
	}
}
