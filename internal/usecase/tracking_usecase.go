package usecase

import (
	"errors"
	"nutrimama/internal/entity"
	"nutrimama/internal/model"
	"nutrimama/internal/repository"

	"gorm.io/gorm"
)

type TrackingUseCase struct {
	DB                   *gorm.DB
	TrackingLogRepo      *repository.TrackingLogRepository
	TrackingQuestionRepo *repository.TrackingQuestionRepository
	MotherRepo           *repository.MotherRepository
}

func NewTrackingUseCase(
	db *gorm.DB, 
	tlRepo *repository.TrackingLogRepository, 
	tqRepo *repository.TrackingQuestionRepository,
	mRepo *repository.MotherRepository,
) *TrackingUseCase {
	return &TrackingUseCase{
		DB:                   db,
		TrackingLogRepo:      tlRepo,
		TrackingQuestionRepo: tqRepo,
		MotherRepo:           mRepo,
	}
}

// =========================================================================
// CRON JOB PLACEHOLDER: NOTIFICATION SYSTEM
// =========================================================================
// TODO: Implement a Cron scheduled job here (e.g., using robfig/cron).
// The Cron job should automatically execute every day/week/month to:
// 1. Query all active Mothers from the database.
// 2. Check the `user_tracking_logs` to see if they submitted their metrics
//    for the current period (Daily, Weekly, or Monthly).
// 3. If missing, trigger a Push Notification / Email reminding them to fill it out!
// 
// Example signature: 
// func (u *TrackingUseCase) NotifyMothersMissingTracking() error { ... }
// =========================================================================

func (u *TrackingUseCase) SubmitTracking(userId int, req model.SubmitTrackingRequest) error {
	c := u.DB.Begin()

	// 1. Lookup the mother profile using the authenticated user_id
	mother, err := u.MotherRepo.FindByUserId(c, userId)
	if err != nil {
		c.Rollback()
		return errors.New("only registered mothers can submit tracking logs")
	}

	// 2. Draft the tracking log entity
	log := entity.UserTrackingLog{
		MotherID:      &mother.MotherID,
		TrackingDate:  req.TrackingDate,
		Frequency:     req.Frequency,
		AnswersData:   req.AnswersData, // JSON payload representing exactly their choices!
		Notes:         req.Notes,
	}
	
	// 3. Save it to unified JSON tracking table
	if err := u.TrackingLogRepo.Create(c, &log); err != nil {
		c.Rollback()
		return err
	}

	return c.Commit().Error
}
