package repository

import (
	"nutrimama/internal/entity"
)

type EduToolsRepository struct {
	Repository[entity.EduTools]
}

func NewEduToolsRepository() *EduToolsRepository {
	return &EduToolsRepository{
		Repository: Repository[entity.EduTools]{},
	}
}
