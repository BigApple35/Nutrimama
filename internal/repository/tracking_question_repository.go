package repository

import (
	"nutrimama/internal/entity"

	"gorm.io/gorm"
)

type TrackingQuestionRepository struct {
	Repository[entity.Question]
}

func NewTrackingQuestionRepository() *TrackingQuestionRepository {
	return &TrackingQuestionRepository{
		Repository: Repository[entity.Question]{},
	}
}

// FindByFrequency fetches all questions that match a particular frequency (daily, weekly, monthly)
func (r *TrackingQuestionRepository) FindByFrequency(db *gorm.DB, frequency string) ([]entity.Question, error) {
	var questions []entity.Question
	// for now this is a placeholder query that fetches them via the generic repository or assumes mapping.
	err := db.Find(&questions).Error
	return questions, err
}
