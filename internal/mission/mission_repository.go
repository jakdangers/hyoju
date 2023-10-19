package mission

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
)

type missionRepository struct {
	gormDB *gorm.DB
}

func NewMissionRepository(gormDB *gorm.DB) *missionRepository {
	return &missionRepository{
		gormDB: gormDB,
	}
}

var _ entity.MissionRepository = (*missionRepository)(nil)

func (m missionRepository) CreateMission(ctx context.Context, mission *entity.Mission) (*entity.Mission, error) {
	const op cerrors.Op = "mission/repository/createMission"

	result := m.gormDB.WithContext(ctx).Create(mission)
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return mission, nil
}

func (m missionRepository) GetMission(ctx context.Context, missionID uint) (*entity.Mission, error) {
	const op cerrors.Op = "mission/repository/getMission"

	var mission entity.Mission
	result := m.gormDB.WithContext(ctx).Where("id = ?", missionID).First(&mission)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, cerrors.E(op, cerrors.Invalid, result.Error)
	}
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return &mission, nil
}

func (m missionRepository) ListMissions(ctx context.Context, userID entity.BinaryUUID) ([]entity.Mission, error) {
	const op cerrors.Op = "mission/repository/listMissions"

	var missions []entity.Mission
	result := m.gormDB.WithContext(ctx).Where("author_id = ? AND status = ?", userID, entity.Active).Find(&missions)
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return missions, nil
}

func (m missionRepository) ListMultipleModeMissions(ctx context.Context, userID entity.BinaryUUID) ([]entity.Mission, error) {
	const op cerrors.Op = "mission/repository/listMultipleModeMissions"

	// TODO 시간 그리고 정렬
	rows, err := m.gormDB.WithContext(ctx).Table("missions").Select(
		"missions.id, missions.author_id, missions.title, missions.emoji, missions.duration, missions.start_date, missions.end_date, missions.plan_date, missions.alarm, missions.week_day, missions.type, missions_status").
		Joins("inner join mission_participants on mission_participants.mission_id = missions.id").
		Where("missions.status = ? AND mission_participants.user_id = ?", entity.Active, userID).Rows()
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	var missions []entity.Mission
	for rows.Next() {
		var mission entity.Mission
		if err := rows.Scan(
			&mission.ID,
			&mission.AuthorID,
			&mission.Title,
			&mission.Emoji,
			&mission.Duration,
			&mission.StartDate,
			&mission.EndDate,
			&mission.PlanTime,
			&mission.Alarm,
			&mission.WeekDay,
			&mission.Type,
			&mission.Status,
		); err != nil {
			return nil, cerrors.E(op, cerrors.Internal, err)
		}
		missions = append(missions, mission)
	}

	return missions, nil
}

func (m missionRepository) ListActiveSingleMissionIDs(ctx context.Context) ([]uint, error) {
	const op cerrors.Op = "mission/repository/listActiveSingleMissionIDs"

	rows, err := m.gormDB.WithContext(ctx).Table("missions").Select("id").Where("type = ? AND status = ?", entity.Single, entity.Active).Rows()
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	var missionIDs []uint
	for rows.Next() {
		var id uint
		err := rows.Scan(&id)
		if err != nil {
			return nil, cerrors.E(op, cerrors.Internal, err)
		}
		missionIDs = append(missionIDs, id)
	}

	return missionIDs, nil
}

func (m missionRepository) PatchMission(ctx context.Context, mission *entity.Mission) (*entity.Mission, error) {
	const op cerrors.Op = "mission/repository/patchMission"

	result := m.gormDB.WithContext(ctx).Save(mission)
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return mission, nil
}
