package challenge_history

import (
	"context"
	"github.com/samber/lo"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
	"time"
)

type challengeHistoryService struct {
	challengeRepo            entity.ChallengeRepository
	challengeParticipantRepo entity.ChallengeParticipantRepository
	challengeHistoryRepo     entity.ChallengeHistoryRepository
	userRepo                 entity.UserRepository
}

func NewChallengeHistoryService(missionRepo entity.ChallengeRepository, missionParticipantRepo entity.ChallengeParticipantRepository, missionHistoryRepo entity.ChallengeHistoryRepository, userRepo entity.UserRepository) *challengeHistoryService {
	return &challengeHistoryService{
		challengeRepo:            missionRepo,
		challengeParticipantRepo: missionParticipantRepo,
		challengeHistoryRepo:     missionHistoryRepo,
		userRepo:                 userRepo,
	}
}

var _ entity.ChallengeHistoryService = (*challengeHistoryService)(nil)

func (m challengeHistoryService) CreateMissionHistory(ctx context.Context, req entity.CreateMissionHistoryRequest) (*entity.CreateMissionHistoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m challengeHistoryService) ListMultiChallengeHistories(ctx context.Context, req entity.ListMultiChallengeHistoriesRequest) (*entity.ListMultiChallengeHistoriesResponse, error) {
	const op cerrors.Op = "challengeHistory/service/ListMultiChallengeHistories"

	// 유저 아이디 파싱
	userID, err := entity.ParseUUID(req.UserID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Invalid, err)
	}

	// 유저 검증
	_, err = m.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	// 입력 날짜를 파싱합니다.
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}
	utcKst := date.Add(-time.Hour * 9)

	// 싱글 또는 멀티 챌린지 조회
	challenges, err := m.challengeRepo.ListMultiChallenges(ctx, entity.ListMultiChallengeParams{
		UserID: userID,
		Date:   utcKst,
		Type:   req.Type,
	})
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	var challengeIDs []uint
	for _, challenge := range challenges {
		challengeIDs = append(challengeIDs, challenge.ID)
	}

	challengeHistories, err := m.challengeHistoryRepo.ListMultiChallengeHistories(ctx, entity.ListMultipleMissionHistoriesParams{
		UserID:       userID,
		ChallengeIDs: challengeIDs,
	})

	var challengeHistoryDTOs []entity.ChallengeHistoryDTO
	for _, challenge := range challenges {
		var historyDTO entity.ChallengeHistoryDTO
		history, ok := lo.Find(challengeHistories, func(item entity.ChallengeHistory) bool {
			return item.ChallengeID == challenge.ID
		})

		historyDTO = entity.ChallengeHistoryDTO{
			ID:          history.ID,
			UserID:      history.UserID.String(),
			ChallengeID: history.ChallengeID,
			Title:       challenge.Title,
			Emoji:       challenge.Emoji,
			PlanTime:    history.PlanTime,
			FrontImage:  history.FrontImage,
			BackImage:   history.BackImage,
		}

		if !ok {
			historyDTO = entity.ChallengeHistoryDTO{
				UserID:      userID.String(),
				ChallengeID: challenge.ID,
				Title:       challenge.Title,
				Emoji:       challenge.Emoji,
				PlanTime:    challenge.PlanTime,
				FrontImage:  "",
				BackImage:   "",
			}
		}

		challengeHistoryDTOs = append(challengeHistoryDTOs, historyDTO)
	}

	return &entity.ListMultiChallengeHistoriesResponse{
		ChallengeHistories: challengeHistoryDTOs,
	}, nil
}
