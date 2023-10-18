package mission_participant

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
	db                           *sql.DB
	gormDB                       *gorm.DB
	sqlMock                      sqlmock.Sqlmock
	missionParticipantRepository entity.MissionParticipantRepository
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
	ts.missionParticipantRepository = NewMissionParticipantRepository(gormDB)

	return &ts
}

func Test_missionParticipantRepository_CreateMissionParticipant(t *testing.T) {
	type args struct {
		ctx         context.Context
		participant *entity.MissionParticipant
	}

	ts := initRepoTestSuite()
	testUserID := entity.BinaryUUIDNew()

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *entity.MissionParticipant
		wantErr bool
	}{
		{
			name: "PASS mission 참여자 생성",
			args: args{
				ctx: context.Background(),
				participant: &entity.MissionParticipant{
					UserID:    testUserID,
					MissionID: 1,
				},
			},
			mock: func() {
				ts.sqlMock.ExpectExec("INSERT INTO `mission_participants`").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: &entity.MissionParticipant{
				Model: gorm.Model{
					ID: 1,
				},
				UserID:    testUserID,
				MissionID: 1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.missionParticipantRepository.CreateMissionParticipant(tt.args.ctx, tt.args.participant)
			assert.Equal(t, true, cmp.Equal(tt.want, got, cmpopts.IgnoreFields(entity.MissionParticipant{}, "CreatedAt", "UpdatedAt", "DeletedAt")))
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_missionParticipantRepository_ListMissionParticipants(t *testing.T) {
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
		want    []entity.MissionParticipant
		wantErr bool
	}{
		{
			name: "PASS mission 참여자 목록 조회",
			args: args{
				ctx:       context.Background(),
				missionID: 1,
			},
			mock: func() {
				query := "SELECT (.+) FROM `mission_participants`"
				columns := []string{"id", "user_id", "mission_id"}
				rows := sqlmock.NewRows(columns).AddRow(1, testUserID, 1)
				ts.sqlMock.ExpectQuery(query).WillReturnRows(rows)
			},
			want: []entity.MissionParticipant{
				{
					Model: gorm.Model{
						ID: 1,
					},
					UserID:    testUserID,
					MissionID: 1,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := ts.missionParticipantRepository.ListMissionParticipants(tt.args.ctx, tt.args.missionID)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}
