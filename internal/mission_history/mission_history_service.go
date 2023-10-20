package mission_history

import (
	"context"
	"github.com/samber/lo"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
	"time"
)

type missionHistoryService struct {
	missionRepo            entity.MissionRepository
	missionParticipantRepo entity.MissionParticipantRepository
	missionHistoryRepo     entity.MissionHistoryRepository
	userRepo               entity.UserRepository
}

func NewMissionHistoryService(missionRepo entity.MissionRepository, missionParticipantRepo entity.MissionParticipantRepository, missionHistoryRepo entity.MissionHistoryRepository, userRepo entity.UserRepository) *missionHistoryService {
	return &missionHistoryService{
		missionRepo:            missionRepo,
		missionParticipantRepo: missionParticipantRepo,
		missionHistoryRepo:     missionHistoryRepo,
		userRepo:               userRepo,
	}
}

var _ entity.MissionHistoryService = (*missionHistoryService)(nil)

func (m missionHistoryService) CreateMissionHistory(ctx context.Context, req entity.CreateMissionHistoryRequest) (*entity.CreateMissionHistoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m missionHistoryService) ListMultiModeMissionHistories(ctx context.Context, req entity.ListMultiModeMissionHistoriesRequest) (*entity.ListMultiModeMissionHistoriesResponse, error) {
	const op cerrors.Op = "missionHistory/service/ListMultipleModeMissionHistories"

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
	// 한국 시간대 (KST) 로드
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}
	kstTime := date.In(loc)
	utcTime := kstTime.UTC()

	// 멀티 플레이 미션 목록 조회
	missions, err := m.missionRepo.ListMultiModeMissions(ctx, entity.ListMultiModeMissionsParams{
		UserID: userID,
		Date:   utcTime,
	})
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	var missionIDs []uint
	for _, mission := range missions {
		missionIDs = append(missionIDs, mission.ID)
	}

	histories, err := m.missionHistoryRepo.ListMultipleModeMissionHistories(ctx, entity.ListMultipleMissionHistoriesParams{
		UserID:     userID,
		MissionIDs: missionIDs,
	})

	var missionHistoryDTOs []entity.MissionHistoryDTO
	for _, mission := range missions {
		var historyDTO entity.MissionHistoryDTO
		history, ok := lo.Find(histories, func(item entity.MissionHistory) bool {
			return item.MissionID == mission.ID
		})

		historyDTO = entity.MissionHistoryDTO{
			ID:         history.ID,
			UserID:     history.UserID.String(),
			MissionID:  history.MissionID,
			Title:      mission.Title,
			Emoji:      mission.Emoji,
			Status:     history.Status,
			PlanTime:   history.PlanTime,
			FrontImage: history.FrontImage,
			BackImage:  history.BackImage,
		}

		if !ok {
			historyDTO = entity.MissionHistoryDTO{
				UserID:     userID.String(),
				MissionID:  mission.ID,
				Title:      mission.Title,
				Emoji:      mission.Emoji,
				Status:     entity.MissionHistoryStatusInit,
				PlanTime:   mission.PlanTime,
				FrontImage: "",
				BackImage:  "",
			}
		}

		missionHistoryDTOs = append(missionHistoryDTOs, historyDTO)
	}

	return &entity.ListMultiModeMissionHistoriesResponse{
		MissionHistories: missionHistoryDTOs,
	}, nil
}
