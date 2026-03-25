package repository

import (
	"nutrimama/internal/entity"

	"gorm.io/gorm"
)

type ConsultantRepository struct {
	Repository[entity.Consultant]
}

func NewConsultantRepository() *ConsultantRepository {
	return &ConsultantRepository{
		Repository: Repository[entity.Consultant]{},
	}
}

func (r *ConsultantRepository) FindByUserId(db *gorm.DB, userId int) (*entity.Consultant, error) {
	var consultant entity.Consultant
	err := db.Where("user_id = ?", userId).First(&consultant).Error
	if err != nil {
		return nil, err
	}
	return &consultant, nil
}
