package single_mode

import (
	"context"
	"fmt"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
)

type singleModeService struct {
	missionRepo            entity.MissionRepository
	missionParticipantRepo entity.MissionParticipantRepository
	missionHistoryRepo     entity.MissionHistoryRepository
	userRepo               entity.UserRepository
}

func NewSingleModeService(missionRepo entity.MissionRepository, missionParticipantRepo entity.MissionParticipantRepository, missionHistoryRepo entity.MissionHistoryRepository, userRepo entity.UserRepository) *singleModeService {
	return &singleModeService{
		missionRepo:            missionRepo,
		missionParticipantRepo: missionParticipantRepo,
		missionHistoryRepo:     missionHistoryRepo,
		userRepo:               userRepo,
	}
}

var _ entity.SingleModeService = (*singleModeService)(nil)

func (s singleModeService) CreateMissionHistories(ctx context.Context) error {
	const op cerrors.Op = "single_mode/service/createMissionHistories"

	// 히스토리 생성 해야 할 전체 미션 아이디 조회
	missionIDs, err := s.missionRepo.ListActiveSingleMissionIDs(ctx)
	if err != nil {
		return cerrors.E(op, cerrors.Internal, err)
	}

	// 미션에 참여한 유저 조회
	for _, missionID := range missionIDs {
		participants, err := s.missionParticipantRepo.ListMissionParticipants(ctx, missionID)
		if err != nil {
			return cerrors.E(op, cerrors.Internal, err)
		}
		// 히스토리 생성
		fmt.Println(participants)
	}

	return nil
}
