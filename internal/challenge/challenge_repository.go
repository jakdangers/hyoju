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
	const op cerrors.Op = "challenge/repository/createChallenge"

	result := m.gormDB.WithContext(ctx).Create(mission)
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return mission, nil
}

func (m challengeRepository) GetChallenge(ctx context.Context, missionID uint) (*entity.Challenge, error) {
	const op cerrors.Op = "challenge/repository/getChallenge"

	var challenge entity.Challenge
	result := m.gormDB.WithContext(ctx).Where("id = ?", missionID).First(&challenge)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, cerrors.E(op, cerrors.Invalid, result.Error)
	}
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return &challenge, nil
}

func (m challengeRepository) ListChallenges(ctx context.Context, params entity.ListChallengesParams) ([]entity.Challenge, error) {
	const op cerrors.Op = "challenge/repository/listChallenges"

	var challenges []entity.Challenge
	result := m.gormDB.WithContext(ctx).Where("user_id = ? AND type = ? AND status = ?", params.UserID, params.Type, entity.ChallengeStatusDeActivate).Find(&challenges)
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return challenges, nil
}

func (m challengeRepository) ListMultiChallenges(ctx context.Context, params entity.ListMultiChallengeParams) ([]entity.Challenge, error) {
	const op cerrors.Op = "challenge/repository/listMultiChallenges"

	rows, err := m.gormDB.WithContext(ctx).Table("challenges").Select(
		"challenges.id, challenges.user_id, challenges.title, challenges.emoji, challenges.duration, challenges.start_date, challenges.end_date, challenges.plan_time, challenges.alarm, challenges.week_day, challenges.type, challenges.status").
		Joins("inner join challenge_participants on challenge_participants.challenge_id = challenges.id").
		Where("challenges.status = ? AND challenge_participants.user_id = ? AND challenge_type AND challenges.start_date <= ? AND challenges.end_date >= ?", entity.ChallengeStatusActivate, params.UserID, params.Type, params.Date, params.Date).Order("challenges.plan_time").Rows()
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	var challenges []entity.Challenge
	for rows.Next() {
		var challenge entity.Challenge
		if err := rows.Scan(
			&challenge.ID,
			&challenge.UserID,
			&challenge.Title,
			&challenge.Emoji,
			&challenge.Duration,
			&challenge.StartDate,
			&challenge.EndDate,
			&challenge.PlanTime,
			&challenge.Alarm,
			&challenge.WeekDay,
			&challenge.Type,
			&challenge.Status,
		); err != nil {
			return nil, cerrors.E(op, cerrors.Internal, err)
		}
		challenges = append(challenges, challenge)
	}

	return challenges, nil
}

func (m challengeRepository) PatchChallenge(ctx context.Context, mission *entity.Challenge) (*entity.Challenge, error) {
	const op cerrors.Op = "challenge/repository/patchMission"

	result := m.gormDB.WithContext(ctx).Save(mission)
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return mission, nil
}

func (m challengeRepository) ChallengeFindByCode(ctx context.Context, code string) (*entity.Challenge, error) {
	const op cerrors.Op = "challenge/repository/challengeFindByCode"

	var challenge entity.Challenge
	result := m.gormDB.WithContext(ctx).Where("code = ?", code).First(&challenge)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return &challenge, nil
}
