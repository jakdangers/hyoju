package group_challenge

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pixelix/entity"
	"testing"
)

type repoTestSuite struct {
	db         *sql.DB
	gormDB     *gorm.DB
	sqlMock    sqlmock.Sqlmock
	repository entity.GroupChallengeRepository
}

func initRepoTestSuite() *repoTestSuite {
	var ts repoTestSuite

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	ts.db = mockDB
	ts.sqlMock = mock

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	ts.gormDB = gormDB
	ts.repository = NewGroupChallengeRepository(gormDB)

	return &ts
}

func Test_groupChallengeRepository_CreateGroupChallenge(t *testing.T) {
	type args struct {
		ctx            context.Context
		groupChallenge *entity.GroupChallenge
	}

	ts := initRepoTestSuite()

	tests := []struct {
		name    string
		mock    func()
		args    args
		want    *entity.GroupChallenge
		wantErr bool
	}{
		{
			name: "PASS 그룹 챌린지 생성",
			mock: func() {
				ts.sqlMock.ExpectExec("INSERT INTO `group_challenges`").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			args: args{
				ctx: context.Background(),
				groupChallenge: &entity.GroupChallenge{
					Title:       "test_group_challenge",
					Description: "test_description",
				},
			},
			want: &entity.GroupChallenge{
				Model: gorm.Model{
					ID: 1,
				},
				Title:       "test_group_challenge",
				Description: "test_description",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.repository.CreateGroupChallenge(tt.args.ctx, tt.args.groupChallenge)
			assert.Equal(t, true, cmp.Equal(tt.want, got, cmpopts.IgnoreFields(gorm.Model{}, "CreatedAt", "UpdatedAt", "DeletedAt")))
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}
