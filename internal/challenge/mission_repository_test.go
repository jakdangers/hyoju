package challenge

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
	db         *sql.DB
	gormDB     *gorm.DB
	sqlMock    sqlmock.Sqlmock
	repository entity.ChallengeRepository
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
	ts.repository = NewChallengeRepository(gormDB)

	return &ts
}

func Test_missionRepository_CreateChallenge(t *testing.T) {
	type args struct {
		ctx       context.Context
		challenge *entity.Challenge
	}

	ts := initRepoTestSuite()
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *entity.Challenge
		wantErr bool
	}{
		{
			name: "PASS 미션 생성",
			args: args{
				ctx: context.Background(),
				challenge: &entity.Challenge{
					UserID:   testUserID,
					Title:    "test_mission",
					Emoji:    "test_emoji",
					Duration: entity.ChallengeDurationPeriod,
					Alarm:    false,
					WeekDay:  3,
					Type:     entity.ChallengeTypeSingle,
					Status:   entity.ChallengeStatusActivate,
				},
			},
			mock: func() {
				ts.sqlMock.ExpectExec("INSERT INTO `challenges` (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: &entity.Challenge{
				Model: gorm.Model{
					ID: 1,
				},
				UserID:   testUserID,
				Title:    "test_mission",
				Emoji:    "test_emoji",
				Duration: entity.ChallengeDurationPeriod,
				Alarm:    false,
				WeekDay:  3,
				Type:     entity.ChallengeTypeSingle,
				Status:   entity.ChallengeStatusActivate,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			got, err := ts.repository.CreateChallenge(tt.args.ctx, tt.args.challenge)
			assert.Equal(t, true, cmp.Equal(tt.want, got, cmpopts.IgnoreFields(entity.Challenge{}, "CreatedAt", "UpdatedAt", "DeletedAt")))
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_missionRepository_ListChallenges(t *testing.T) {
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
		want    []entity.Challenge
		wantErr bool
	}{
		{
			name: "PASS 미션 목록 조회",
			args: args{
				ctx:    context.Background(),
				userID: testUserID,
			},
			mock: func() {
				query := "SELECT (.+) FROM `challenges`"
				columns := []string{"id", "user_id", "title", "emoji", "duration", "alarm", "week_day", "type", "status"}
				rows := sqlmock.NewRows(columns).AddRow(1, testUserID, "test_mission", "test_emoji", "DAILY", true, 3, "SINGLE", "ACTIVATE")
				ts.sqlMock.ExpectQuery(query).WillReturnRows(rows)
			},
			want: []entity.Challenge{
				{
					Model: gorm.Model{
						ID: 1,
					},
					UserID:   testUserID,
					Title:    "test_mission",
					Emoji:    "test_emoji",
					Duration: entity.ChallengeDurationDaily,
					Alarm:    true,
					WeekDay:  3,
					Type:     entity.ChallengeTypeSingle,
					Status:   entity.ChallengeStatusActivate,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.repository.ListChallenges(tt.args.ctx, tt.args.userID)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_missionRepository_PatchMission(t *testing.T) {
	type args struct {
		ctx       context.Context
		challenge *entity.Challenge
	}

	ts := initRepoTestSuite()
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *entity.Challenge
		wantErr bool
	}{
		{
			name: "PASS 첼린지 수정",
			args: args{
				ctx: context.Background(),
				challenge: &entity.Challenge{
					Model: gorm.Model{
						ID: 1,
					},
					UserID:   testUserID,
					Title:    "modified_mission",
					Emoji:    "modified_emoji",
					Duration: entity.ChallengeDurationPeriod,
					Alarm:    false,
					WeekDay:  7,
					Type:     entity.ChallengeTypeSingle,
					Status:   entity.ChallengeStatusActivate,
				},
			},
			mock: func() {
				query := "UPDATE `challenges`"
				ts.sqlMock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: &entity.Challenge{
				Model: gorm.Model{
					ID: 1,
				},
				UserID:   testUserID,
				Title:    "modified_mission",
				Emoji:    "modified_emoji",
				Duration: entity.ChallengeDurationPeriod,
				Alarm:    false,
				WeekDay:  7,
				Type:     entity.ChallengeTypeSingle,
				Status:   entity.ChallengeStatusActivate,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.repository.PatchChallenge(tt.args.ctx, tt.args.challenge)
			assert.Equal(t, true, cmp.Equal(tt.want, got, cmpopts.IgnoreFields(entity.Challenge{}, "CreatedAt", "UpdatedAt", "DeletedAt")))
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_missionRepository_GetChallenge(t *testing.T) {
	type args struct {
		ctx         context.Context
		challengeID uint
	}

	ts := initRepoTestSuite()
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *entity.Challenge
		wantErr bool
	}{
		{
			name: "PASS challenge 조회",
			args: args{
				ctx:         context.Background(),
				challengeID: 1,
			},
			mock: func() {
				query := "SELECT (.+) FROM `challenges`"
				columns := []string{"id", "user_id", "title", "emoji", "duration", "alarm", "week_day", "type", "status"}
				rows := sqlmock.NewRows(columns).AddRow(1, testUserID, "test_mission", "test_emoji", "DAILY", true, 3, "SINGLE", "ACTIVATE")
				ts.sqlMock.ExpectQuery(query).WillReturnRows(rows)
			},
			want: &entity.Challenge{
				Model: gorm.Model{
					ID: 1,
				},
				UserID:   testUserID,
				Title:    "test_mission",
				Emoji:    "test_emoji",
				Duration: "DAILY",
				Alarm:    true,
				WeekDay:  3,
				Type:     entity.ChallengeTypeSingle,
				Status:   entity.ChallengeStatusActivate,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.repository.GetChallenge(tt.args.ctx, tt.args.challengeID)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_missionRepository_ListMultiChallenges(t *testing.T) {
	type args struct {
		ctx    context.Context
		params entity.ListMultiChallengeParams
	}

	ts := initRepoTestSuite()
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    []entity.Challenge
		wantErr bool
	}{
		{
			name: "PASS challenge 목록 조회",
			args: args{
				ctx: context.Background(),
				params: entity.ListMultiChallengeParams{
					UserID: testUserID,
					Date:   time.Time{},
				},
			},
			mock: func() {
				query := "SELECT (.+) FROM `challenges`"
				columns := []string{"challenges.id", "challenges.user_id", "challenges.title", "challenges.emoji", "challenges.duration", "challenges.start_date", "challenges.end_date", "challenges.plan_date", "challenges.alarm", "challenges.week_day", "challenges.type", "challenges_status"}
				rows := sqlmock.NewRows(columns).AddRow(1, testUserID, "test_mission", "test_emoji", "DAILY", time.Time{}, time.Time{}, 0, true, 3, "SINGLE", "ACTIVATE")
				ts.sqlMock.ExpectQuery(query).WillReturnRows(rows)
			},
			want: []entity.Challenge{
				{
					Model: gorm.Model{
						ID: 1,
					},
					UserID:   testUserID,
					Title:    "test_mission",
					Emoji:    "test_emoji",
					Duration: "DAILY",
					Alarm:    true,
					WeekDay:  3,
					Type:     "SINGLE",
					Status:   "ACTIVE",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.repository.ListMultiChallenges(tt.args.ctx, tt.args.params)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}
