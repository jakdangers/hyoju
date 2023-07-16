package repository

import (
	"context"
	"cryptoChallenges/internal/user/entity"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

type userRepositoryTestSuite struct {
	suite.Suite
	gormDB  *gorm.DB
	sqlxDB  *sqlx.DB
	ctx     context.Context
	sqlMock sqlmock.Sqlmock
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(userRepositoryTestSuite))
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func (us *userRepositoryTestSuite) SetupSuite() {
	// mockDB init
	sqlDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	us.NoError(err)
	// gormDB init
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	us.NoError(err)
	// dependency injection
	us.gormDB = gormDB
	us.sqlMock = mock
	us.ctx = context.Background()
}

func (us *userRepositoryTestSuite) Test_userRepository_CreateUser() {
	tests := map[string]struct {
		input   *entity.User
		want    *entity.User
		mock    func()
		wantErr bool
	}{
		"success": {
			input: &entity.User{
				Name:     "cryptoChallenges",
				Email:    "cryptoChallenges@gmail.com",
				UserID:   "cryptoChallenges",
				Password: "password",
			},
			want: &entity.User{
				Name:     "cryptoChallenges",
				Email:    "cryptoChallenges@gmail.com",
				UserID:   "cryptoChallenges",
				Password: "password",
			},
			mock: func() {
				us.sqlMock.ExpectExec("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`email`,`user_id`,`password`) VALUES (?,?,?,?,?,?,?)").
					WithArgs(
						AnyTime{},
						AnyTime{},
						nil,
						"cryptoChallenges",
						"cryptoChallenges@gmail.com",
						"cryptoChallenges",
						"password",
					).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		ur := New(us.gormDB, us.sqlxDB)
		test.mock()
		got, err := ur.CreateUser(us.ctx, test.want)
		if us.NoError(err) {
			isEqual := cmp.Equal(test.want, got, cmpopts.IgnoreFields(entity.User{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"))
			us.Equal(true, isEqual)
		}
		//us.EqualErrorf(err, expectedErrorString, "error message %s", "formatted")
		//us.Errorf(err, "error message %s", "formatted")
		if us.Error(err) {

		}
	}
}
