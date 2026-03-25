package repository

import (
	"nutrimama/internal/entity"
)

type TrackingLogRepository struct {
	Repository[entity.UserTrackingLog]
}

func NewTrackingLogRepository() *TrackingLogRepository {
	return &TrackingLogRepository{
		Repository: Repository[entity.UserTrackingLog]{},
	}
}
