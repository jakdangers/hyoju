package challenge_history

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pixelix/entity"
	"testing"
	"time"
)

type repoTestSuite struct {
	db         *sql.DB
	gormDB     *gorm.DB
	sqlMock    sqlmock.Sqlmock
	repository entity.ChallengeHistoryRepository
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
	ts.repository = NewChallengeHistoryRepository(gormDB)

	return &ts
}

func Test_missionHistoryRepository_ListMultipleModeMissionHistories(t *testing.T) {
	type args struct {
		ctx    context.Context
		params entity.ListGroupChallengeHistoriesParams
	}

	ts := initRepoTestSuite()
	testUserID := entity.BinaryUUIDNew()
	testTimeStamp := time.Date(2023, 10, 10, 10, 10, 10, 00, time.UTC)

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    []entity.ChallengeHistory
		wantErr bool
	}{
		{
			name: "PASS challenge history 리스트 조회",
			args: args{
				ctx: context.Background(),
				params: entity.ListGroupChallengeHistoriesParams{
					UserID:       testUserID,
					ChallengeIDs: []uint{1, 2, 3},
				},
			},

			mock: func() {
				query := "SELECT (.+) FROM `mission_histories`"
				columns := []string{"id", "user_id", "mission_id", "status", "date", "plan_time", "front_image", "back_image"}
				rows := sqlmock.NewRows(columns).AddRow(1, testUserID, 1, "INIT", testTimeStamp, testTimeStamp, "", "")
				ts.sqlMock.ExpectQuery(query).WillReturnRows(rows)
			},
			want: []entity.ChallengeHistory{
				{
					Model: gorm.Model{
						ID: 1,
					},
					UserID:      testUserID,
					ChallengeID: 1,
					PlanTime:    testTimeStamp,
					FrontImage:  "",
					BackImage:   "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.repository.ListGroupChallengeHistories(tt.args.ctx, tt.args.params)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}
