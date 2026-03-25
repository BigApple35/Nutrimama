package repository

import (
	"nutrimama/internal/entity"

	"gorm.io/gorm"
)

type MotherRepository struct {
	Repository[entity.Mother]
}

func NewMotherRepository() *MotherRepository {
	return &MotherRepository{
		Repository: Repository[entity.Mother]{},
	}
}

func (r *MotherRepository) FindByUserId(db *gorm.DB, userId int) (*entity.Mother, error) {
	var mother entity.Mother
	err := db.Where("user_id = ?", userId).First(&mother).Error
	if err != nil {
		return nil, err
	}
	return &mother, nil
}
