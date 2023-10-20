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
					Duration: entity.Period,
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
				Duration:  entity.Period,
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

func Test_missionRepository_PatchMission(t *testing.T) {
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
			name: "PASS 미션 수정",
			args: args{
				ctx: context.Background(),
				mission: &entity.Mission{
					Model: gorm.Model{
						ID: 1,
					},
					AuthorID: testUserID,
					Title:    "modified_mission",
					Emoji:    "modified_emoji",
					Duration: entity.Period,
					Alarm:    false,
					WeekDay:  7,
					Type:     entity.Single,
					Status:   entity.Active,
				},
			},
			mock: func() {
				query := "UPDATE `missions`"
				ts.sqlMock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: &entity.Mission{
				Model: gorm.Model{
					ID: 1,
				},
				AuthorID: testUserID,
				Title:    "modified_mission",
				Emoji:    "modified_emoji",
				Duration: entity.Period,
				Alarm:    false,
				WeekDay:  7,
				Type:     entity.Single,
				Status:   entity.Active,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.missionRepository.PatchMission(tt.args.ctx, tt.args.mission)
			assert.Equal(t, true, cmp.Equal(tt.want, got, cmpopts.IgnoreFields(entity.Mission{}, "CreatedAt", "UpdatedAt", "DeletedAt")))
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_missionRepository_GetMission(t *testing.T) {
	type args struct {
		ctx       context.Context
		missionID uint
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
			name: "PASS 미션 조회",
			args: args{
				ctx:       context.Background(),
				missionID: 1,
			},
			mock: func() {
				query := "SELECT (.+) FROM `missions`"
				columns := []string{"id", "author_id", "title", "emoji", "duration", "alarm", "week_day", "type", "status"}
				rows := sqlmock.NewRows(columns).AddRow(1, testUserID, "test_mission", "test_emoji", "DAILY", true, 3, "SINGLE", "ACTIVE")
				ts.sqlMock.ExpectQuery(query).WillReturnRows(rows)
			},
			want: &entity.Mission{
				Model: gorm.Model{
					ID: 1,
				},
				AuthorID: testUserID,
				Title:    "test_mission",
				Emoji:    "test_emoji",
				Duration: "DAILY",
				Alarm:    true,
				WeekDay:  3,
				Type:     "SINGLE",
				Status:   "ACTIVE",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.missionRepository.GetMission(tt.args.ctx, tt.args.missionID)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_missionRepository_ListActiveSingleMissionIDs(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	ts := initRepoTestSuite()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    []uint
		wantErr bool
	}{
		{
			name: "PASS missionID 리스트 조회",
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				query := "SELECT (.+) FROM `missions`"
				columns := []string{"id"}
				rows := sqlmock.NewRows(columns).AddRow(1).AddRow(2).AddRow(3)
				ts.sqlMock.ExpectQuery(query).WillReturnRows(rows)
			},
			want:    []uint{1, 2, 3},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.missionRepository.ListActiveSingleMissionIDs(tt.args.ctx)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_missionRepository_ListMultiModeMissions(t *testing.T) {
	type args struct {
		ctx    context.Context
		params entity.ListMultiModeMissionsParams
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
				ctx: context.Background(),
				params: entity.ListMultiModeMissionsParams{
					UserID: testUserID,
					Date:   time.Time{},
				},
			},
			mock: func() {
				query := "SELECT (.+) FROM `missions`"
				columns := []string{"missions.id", "missions.author_id", "missions.title", "missions.emoji", "missions.duration", "missions.start_date", "missions.end_date", "missions.plan_date", "missions.alarm", "missions.week_day", "missions.type", "missions_status"}
				rows := sqlmock.NewRows(columns).AddRow(1, testUserID, "test_mission", "test_emoji", "DAILY", time.Time{}, time.Time{}, time.Time{}, true, 3, "SINGLE", "ACTIVE")
				ts.sqlMock.ExpectQuery(query).WillReturnRows(rows)
			},
			want: []entity.Mission{
				{
					Model: gorm.Model{
						ID: 1,
					},
					AuthorID:  testUserID,
					Title:     "test_mission",
					Emoji:     "test_emoji",
					Duration:  "DAILY",
					StartDate: time.Time{},
					EndDate:   time.Time{},
					PlanTime:  time.Time{},
					Alarm:     true,
					WeekDay:   3,
					Type:      "SINGLE",
					Status:    "ACTIVE",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.missionRepository.ListMultiModeMissions(tt.args.ctx, tt.args.params)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}
