package user

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pixelix/entity"
	"strings"
	"testing"
)

type repoTestSuite struct {
	db             *sql.DB
	gormDB         *gorm.DB
	sqlMock        sqlmock.Sqlmock
	userRepository entity.UserRepository
}

func initRepoTestSuite() *repoTestSuite {
	var us repoTestSuite

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	us.db = mockDB
	us.sqlMock = mock

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	us.gormDB = gormDB
	us.userRepository = NewUserRepository(gormDB)

	return &us
}

func Test_userRepository_CreateUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *entity.User
	}

	us := initRepoTestSuite()
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *entity.User
		wantErr bool
	}{
		{
			name: "PASS 유저 생성",
			args: args{
				ctx: context.Background(),
				user: &entity.User{
					Base:        entity.Base{ID: testUserID},
					NickName:    "blipix",
					Email:       "blipix@blipix.com",
					Provider:    "blipix",
					FirebaseUID: "firebaseUID",
				},
			},
			mock: func() {
				us.sqlMock.ExpectExec("INSERT INTO `users`").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: &entity.User{
				Base:        entity.Base{ID: testUserID},
				NickName:    "blipix",
				Email:       "blipix@blipix.com",
				Provider:    "blipix",
				FirebaseUID: "firebaseUID",
			},
			wantErr: false,
		},
		{
			name: "FAIL 중복 email 유저 생성",
			args: args{
				ctx: context.Background(),
				user: &entity.User{
					Base:        entity.Base{ID: testUserID},
					NickName:    "blipix",
					Email:       "blipix@blipix.com",
					Provider:    "blipix",
					FirebaseUID: "firebaseUID",
				},
			},
			mock: func() {
				us.sqlMock.ExpectExec("INSERT INTO `users` (.+)").WillReturnError(gorm.ErrDuplicatedKey)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := us.userRepository.CreateUser(tt.args.ctx, tt.args.user)
			assert.Equal(t, true, cmp.Equal(tt.want, got, cmpopts.IgnoreFields(entity.User{}, "CreatedAt", "UpdatedAt", "DeletedAt")))
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_userRepository_UpdateUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *entity.User
	}

	us := initRepoTestSuite()
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *entity.User
		wantErr bool
	}{
		{
			name: "PASS 존재하는 userID로 수정 ",
			args: args{
				ctx: context.Background(),
				user: &entity.User{
					Base: entity.Base{
						ID: testUserID,
					},
					NickName:    "modifed_nickName",
					Email:       "original_email",
					Provider:    "blipix",
					FirebaseUID: "firebaseUID",
				},
			},
			mock: func() {
				us.sqlMock.ExpectExec("UPDATE `users`").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: &entity.User{
				Base: entity.Base{
					ID: testUserID,
				},
				NickName:    "modifed_nickName",
				Email:       "original_email",
				Provider:    "blipix",
				FirebaseUID: "firebaseUID",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := us.userRepository.UpdateUser(tt.args.ctx, tt.args.user)
			assert.Equal(t, true, cmp.Equal(tt.want, got, cmpopts.IgnoreFields(entity.User{}, "CreatedAt", "UpdatedAt", "DeletedAt")))
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_userRepository_DeleteUser(t *testing.T) {
	type args struct {
		ctx context.Context
		id  entity.BinaryUUID
	}

	us := initRepoTestSuite()

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
				id:  entity.BinaryUUIDNew(),
			},
			mock: func() {
				us.sqlMock.ExpectExec("UPDATE `users`").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := us.userRepository.DeleteUser(tt.args.ctx, tt.args.id)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_userRepository_FindByEmail(t *testing.T) {
	type args struct {
		ctx   context.Context
		email string
	}

	us := initRepoTestSuite()
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *entity.User
		wantErr bool
	}{
		{
			name: "PASS 올바른 이메일로 조회",
			args: args{
				ctx:   context.Background(),
				email: "blipix@blipix.com",
			},
			mock: func() {
				query := "SELECT (.+) FROM `users` (.+)"
				columns := []string{"id", "nick_name", "email", "provider", "firebase_uid", "friend_code"}
				rows := sqlmock.NewRows(columns).AddRow(
					testUserID, "nick_name", "blipix@blipix.com", "blipix", "firebase_uid", "test_friendCode",
				)
				us.sqlMock.ExpectQuery(query).WillReturnRows(rows)
			},
			want: &entity.User{
				Base: entity.Base{
					ID: testUserID,
				},
				NickName:    "nick_name",
				Email:       "blipix@blipix.com",
				Provider:    "blipix",
				FirebaseUID: "firebase_uid",
				FriendCode:  "test_friendCode",
			},
			wantErr: false,
		},
		{
			name: "PASS 존재 하지 않는 이메일 조회",
			args: args{
				ctx:   context.Background(),
				email: "blipix@blipix.com",
			},
			mock: func() {
				query := "SELECT (.+) FROM `users` (.+)"
				us.sqlMock.ExpectQuery(query).WillReturnError(gorm.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := us.userRepository.FindByEmail(tt.args.ctx, tt.args.email)
			assert.Equal(t, true, cmp.Equal(tt.want, got, cmpopts.IgnoreFields(entity.User{}, "CreatedAt", "UpdatedAt", "DeletedAt")))
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_userRepository_FindByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  entity.BinaryUUID
	}

	us := initRepoTestSuite()
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *entity.User
		wantErr bool
	}{
		{
			name: "PASS 존재하는 userID로 조회",
			args: args{
				ctx: context.Background(),
				id:  testUserID,
			},
			mock: func() {
				query := "SELECT (.+) FROM `users` (.+)"
				columns := []string{"id", "nick_name", "email", "provider", "firebase_uid"}
				rows := sqlmock.NewRows(columns).AddRow(
					testUserID, "nick_name", "blipix@blipix.com", "blipix", "firebase_uid",
				)
				us.sqlMock.ExpectQuery(query).WillReturnRows(rows)
			},
			want: &entity.User{
				Base: entity.Base{
					ID: testUserID,
				},
				NickName:    "nick_name",
				Email:       "blipix@blipix.com",
				Provider:    "blipix",
				FirebaseUID: "firebase_uid",
			},
			wantErr: false,
		},
		{
			name: "PASS 존재 하지 않는 userID로 조회",
			args: args{
				ctx: context.Background(),
				id:  testUserID,
			},
			mock: func() {
				query := "SELECT (.+) FROM `users` (.+)"
				us.sqlMock.ExpectQuery(query).WillReturnError(gorm.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := us.userRepository.FindByID(tt.args.ctx, tt.args.id)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_userRepository_FindByFriendCode(t *testing.T) {
	type args struct {
		ctx        context.Context
		friendCode string
	}

	ts := initRepoTestSuite()
	testUserID := entity.BinaryUUIDNew()
	testFriendCode := strings.Split(testUserID.String(), "-")[0]

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *entity.User
		wantErr bool
	}{
		{
			name: "PASS 존재하는 friendCode로 조회",
			args: args{
				ctx:        context.Background(),
				friendCode: testFriendCode,
			},
			mock: func() {
				query := "SELECT (.+) FROM `users`"
				columns := []string{"id", "nick_name", "email", "provider", "firebase_uid", "friend_code"}
				rows := sqlmock.NewRows(columns).AddRow(testUserID, "test_nickName", "blipix@blipix.com", "kakao", "test_firebaseUID", testFriendCode)
				ts.sqlMock.ExpectQuery(query).WillReturnRows(rows)
			},
			want: &entity.User{
				Base: entity.Base{
					ID: testUserID,
				},
				NickName:    "test_nickName",
				Email:       "blipix@blipix.com",
				Provider:    "kakao",
				FirebaseUID: "test_firebaseUID",
				FriendCode:  testFriendCode,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.userRepository.FindByFriendCode(tt.args.ctx, tt.args.friendCode)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}
