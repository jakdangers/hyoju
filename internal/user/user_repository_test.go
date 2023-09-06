package user

import (
	"context"
	"cryptoChallenges/entity"
	"database/sql"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

type userRepositoryTestSuite struct {
	suite.Suite
	db             *sql.DB
	gormDB         *gorm.DB
	sqlMock        sqlmock.Sqlmock
	userRepository entity.UserRepository
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
	mockDB, mock, err := sqlmock.New()
	us.NoError(err)
	us.db = mockDB
	us.sqlMock = mock

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	us.NoError(err)

	us.gormDB = gormDB
	us.userRepository = NewUserRepository(gormDB)
}

func (us *userRepositoryTestSuite) AfterTest(suiteName, testName string) {
	_ = us.db.Close()
}

func (us *userRepositoryTestSuite) Test_userRepository_CreateUser() {
	id := uuid.New()
	tests := []struct {
		name    string
		input   *entity.User
		ctx     context.Context
		mock    func()
		want    *entity.User
		wantErr bool
	}{
		{
			name: "성공-기본",
			input: &entity.User{
				Base: entity.Base{
					ID: id,
				},
				Name:     "cryptoChallenges",
				Email:    "cryptoChallenges@gmail.com",
				UserID:   "cryptoChallenges",
				Password: "password",
			},
			ctx: context.Background(),
			mock: func() {
				us.sqlMock.ExpectExec("INSERT INTO `users`").
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
			want: &entity.User{
				Base: entity.Base{
					ID: id,
				},
				Name:     "cryptoChallenges",
				Email:    "cryptoChallenges@gmail.com",
				UserID:   "cryptoChallenges",
				Password: "password",
			},
			wantErr: false,
		},
		{
			name: "실패-아이디중복",
			input: &entity.User{
				Base: entity.Base{
					ID: id,
				},
				Name:     "cryptoChallenges",
				Email:    "cryptoChallenges@gmail.com",
				UserID:   "cryptoChallenges",
				Password: "password",
			},
			mock: func() {
				us.sqlMock.ExpectExec("INSERT INTO `users`").
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
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		us.T().Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := us.userRepository.CreateUser(tt.ctx, tt.want)
			us.Equal(true, cmp.Equal(tt.want, got, cmpopts.IgnoreFields(entity.User{}, "CreatedAt", "UpdatedAt", "DeletedAt")))
			us.Equal(tt.wantErr, err != nil)
		})
	}
}

func (us *userRepositoryTestSuite) Test_userRepository_ReadUser() {
	testUserID := uuid.New()

	tests := []struct {
		name    string
		input   *entity.User
		ctx     context.Context
		mock    func()
		want    *entity.User
		wantErr bool
	}{
		{
			name: "PASS 기본 userID",
			input: &entity.User{
				Base: entity.Base{
					ID: testUserID,
				},
			},
			ctx: context.Background(),
			mock: func() {
				query := "SELECT (.+) FROM `users`"
				columns := []string{"id", "name"}
				rows := sqlmock.NewRows(columns).AddRow(
					testUserID, "cryptoChallenges",
				)
				us.sqlMock.ExpectQuery(query).WillReturnRows(rows)
			},
			want: &entity.User{
				Base: entity.Base{
					ID: testUserID,
				},
				Name: "cryptoChallenges",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		us.T().Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := us.userRepository.ReadUser(tt.ctx, tt.input)
			us.Equal(true, cmp.Equal(tt.want, got, cmpopts.IgnoreFields(entity.User{}, "CreatedAt", "UpdatedAt", "DeletedAt")))
			us.Equal(tt.wantErr, err != nil)
		})
	}
}
