package challenge

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
)

type challengeRepository struct {
	gormDB *gorm.DB
}

func NewChallengeRepository(gormDB *gorm.DB) *challengeRepository {
	return &challengeRepository{
		gormDB: gormDB,
	}
}

var _ entity.ChallengeRepository = (*challengeRepository)(nil)

func (m challengeRepository) CreateChallenge(ctx context.Context, mission *entity.Challenge) (*entity.Challenge, error) {
	const op cerrors.Op = "challenge/repository/createMission"

	result := m.gormDB.WithContext(ctx).Create(mission)
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return mission, nil
}

func (m challengeRepository) GetChallenge(ctx context.Context, missionID uint) (*entity.Challenge, error) {
	const op cerrors.Op = "challenge/repository/getMission"

	var mission entity.Challenge
	result := m.gormDB.WithContext(ctx).Where("id = ?", missionID).First(&mission)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, cerrors.E(op, cerrors.Invalid, result.Error)
	}
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return &mission, nil
}

func (m challengeRepository) ListChallenges(ctx context.Context, userID entity.BinaryUUID) ([]entity.Challenge, error) {
	const op cerrors.Op = "challenge/repository/listMissions"

	var missions []entity.Challenge
	result := m.gormDB.WithContext(ctx).Where("user_id = ? AND status = ?", userID, entity.ChallengeStatusDeActivate).Find(&missions)
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return missions, nil
}

func (m challengeRepository) ListMultiChallenges(ctx context.Context, params entity.ListMultiChallengeParams) ([]entity.Challenge, error) {
	const op cerrors.Op = "challenge/repository/listMultipleModeMissions"

	rows, err := m.gormDB.WithContext(ctx).Table("missions").Select(
		"missions.id, missions.user_id, missions.title, missions.emoji, missions.duration, missions.start_date, missions.end_date, missions.plan_time, missions.alarm, missions.week_day, missions.type, missions.status").
		Joins("inner join mission_participants on mission_participants.mission_id = missions.id").
		Where("missions.status = ? AND mission_participants.user_id = ? AND missions.start_date <= ? AND missions.end_date >= ?", entity.ChallengeStatusActivate, params.UserID, params.Date, params.Date).Order("missions.plan_time").Rows()
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	var missions []entity.Challenge
	for rows.Next() {
		var mission entity.Challenge
		if err := rows.Scan(
			&mission.ID,
			&mission.UserID,
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

func (m challengeRepository) PatchChallenge(ctx context.Context, mission *entity.Challenge) (*entity.Challenge, error) {
	const op cerrors.Op = "challenge/repository/patchMission"

	result := m.gormDB.WithContext(ctx).Save(mission)
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return mission, nil
}
