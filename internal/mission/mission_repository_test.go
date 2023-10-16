package mission

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
	"time"
)

type repoTestSuite struct {
	db                *sql.DB
	gormDB            *gorm.DB
	sqlMock           sqlmock.Sqlmock
	missionRepository entity.MissionRepository
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
	ts.missionRepository = NewMissionRepository(gormDB)

	return &ts
}

func Test_missionRepository_CreateMission(t *testing.T) {
	type args struct {
		ctx     context.Context
		mission *entity.Mission
	}

	ts := initRepoTestSuite()
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *entity.Mission
		wantErr bool
	}{
		{
			name: "PASS 미션 생성",
			args: args{
				ctx: context.Background(),
				mission: &entity.Mission{
					AuthorID: testUserID,
					Title:    "test_mission",
					Emoji:    "test_emoji",
					Duration: entity.WeekDay,
					Alarm:    false,
					WeekDay:  3,
					Type:     entity.Single,
					Status:   entity.Active,
				},
			},
			mock: func() {
				ts.sqlMock.ExpectExec("INSERT INTO `missions` (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: &entity.Mission{
				Model: gorm.Model{
					ID: 1,
				},
				AuthorID:  testUserID,
				Title:     "test_mission",
				Emoji:     "test_emoji",
				Duration:  entity.WeekDay,
				StartDate: time.Time{},
				EndDate:   time.Time{},
				PlanTime:  time.Time{},
				Alarm:     false,
				WeekDay:   3,
				Type:      entity.Single,
				Status:    entity.Active,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			got, err := ts.missionRepository.CreateMission(tt.args.ctx, tt.args.mission)
			assert.Equal(t, true, cmp.Equal(tt.want, got, cmpopts.IgnoreFields(entity.Mission{}, "CreatedAt", "UpdatedAt", "DeletedAt")))
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_missionRepository_ListMissions(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID entity.BinaryUUID
	}

	ts := initRepoTestSuite()
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    []entity.Mission
		wantErr bool
	}{
		{
			name: "PASS 미션 목록 조회",
			args: args{
				ctx:    context.Background(),
				userID: testUserID,
			},
			mock: func() {
				query := "SELECT (.+) FROM `missions`"
				columns := []string{"id", "author_id", "title", "emoji", "duration", "alarm", "week_day", "type", "status"}
				rows := sqlmock.NewRows(columns).AddRow(1, testUserID, "test_mission", "test_emoji", "DAILY", true, 3, "SINGLE", "ACTIVE")
				ts.sqlMock.ExpectQuery(query).WillReturnRows(rows)
			},
			want: []entity.Mission{
				{
					Model: gorm.Model{
						ID: 1,
					},
					AuthorID: testUserID,
					Title:    "test_mission",
					Emoji:    "test_emoji",
					Duration: entity.Daily,
					Alarm:    true,
					WeekDay:  3,
					Type:     entity.Single,
					Status:   entity.Active,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.missionRepository.ListMissions(tt.args.ctx, tt.args.userID)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}
