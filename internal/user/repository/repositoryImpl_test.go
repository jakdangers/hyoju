package repository

import (
	"context"
	"cryptoChallenges/internal/user/entity"
	"cryptoChallenges/pkg/models"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

type userRepositoryTestSuite struct {
	suite.Suite
	gormDB         *gorm.DB
	sqlxDB         *sqlx.DB
	ctx            context.Context
	sqlMock        sqlmock.Sqlmock
	userRepository UserRepository
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(userRepositoryTestSuite))
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func (us *userRepositoryTestSuite) SetupTest() {
	// dependency init
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	us.NoError(err)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	us.NoError(err)

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// dependency injection
	us.NoError(err)
	us.gormDB = gormDB
	us.sqlMock = mock
	us.ctx = context.Background()
	us.userRepository = New(gormDB, sqlxDB)
}

func (us *userRepositoryTestSuite) Test_userRepository_CreateUser() {
	id := uuid.New()
	tests := []struct {
		name    string
		given   *entity.User
		want    *entity.User
		mock    func()
		wantErr bool
	}{
		{
			name: "성공-기본",
			given: &entity.User{
				Base: models.Base{
					ID: id,
				},
				Name:     "cryptoChallenges",
				Email:    "cryptoChallenges@gmail.com",
				UserID:   "cryptoChallenges",
				Password: "password",
			},
			want: &entity.User{
				Base: models.Base{
					ID: id,
				},
				Name:     "cryptoChallenges",
				Email:    "cryptoChallenges@gmail.com",
				UserID:   "cryptoChallenges",
				Password: "password",
			},
			mock: func() {
				us.sqlMock.ExpectExec("INSERT INTO `users` (`id`,`created_at`,`updated_at`,`deleted_at`,`name`,`email`,`user_id`,`password`) VALUES (?,?,?,?,?,?,?,?)").
					WithArgs(
						id,
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
		{
			name: "실패-아이디중복",
			given: &entity.User{
				Base: models.Base{
					ID: id,
				},
				Name:     "cryptoChallenges",
				Email:    "cryptoChallenges@gmail.com",
				UserID:   "cryptoChallenges",
				Password: "password",
			},
			want: &entity.User{
				Base: models.Base{
					ID: id,
				},
				Name:     "cryptoChallenges",
				Email:    "cryptoChallenges@gmail.com",
				UserID:   "cryptoChallenges",
				Password: "password",
			},
			mock: func() {
				us.sqlMock.ExpectExec("INSERT INTO `users` (`id`,`created_at`,`updated_at`,`deleted_at`,`name`,`email`,`user_id`,`password`) VALUES (?,?,?,?,?,?,?,?)").
					WithArgs(
						id,
						AnyTime{},
						AnyTime{},
						nil,
						"cryptoChallenges",
						"cryptoChallenges@gmail.com",
						"cryptoChallenges",
						"password",
					).WillReturnError(gorm.ErrDuplicatedKey)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		us.T().Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := us.userRepository.CreateUser(us.ctx, tt.want)
			if err == nil {
				us.Equal(true, cmp.Equal(tt.want, got, cmpopts.IgnoreFields(entity.User{}, "CreatedAt", "UpdatedAt", "DeletedAt")))
			}
			if err != nil {
				us.EqualError(err, "user/createUser: internal error: duplicated key not allowed")
			}
		})
	}
}
