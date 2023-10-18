package mission_history

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"pixelix/entity"
	"testing"
)

type repoTestSuite struct {
	db         *sql.DB
	gormDB     *gorm.DB
	sqlMock    sqlmock.Sqlmock
	repository entity.MissionHistoryRepository
}

func initRepo

func Test_missionHistoryRepository_CreateMissionHistory(t *testing.T) {
	type fields struct {
		gormDB *gorm.DB
	}
	type args struct {
		ctx            context.Context
		missionHistory *entity.MissionHistory
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.MissionHistory
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := missionHistoryRepository{
				gormDB: tt.fields.gormDB,
			}
			got, err := m.CreateMissionHistory(tt.args.ctx, tt.args.missionHistory)
			if !tt.wantErr(t, err, fmt.Sprintf("CreateMissionHistory(%v, %v)", tt.args.ctx, tt.args.missionHistory)) {
				return
			}
			assert.Equalf(t, tt.want, got, "CreateMissionHistory(%v, %v)", tt.args.ctx, tt.args.missionHistory)
		})
	}
}
