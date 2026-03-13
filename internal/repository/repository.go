package repository

import gorm "gorm.io/gorm"

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *Repository[T]) List() ([]T, error) {
	var entities []T
	err := r.DB.Find(&entities).Error
	return entities, err
	
}

func (r *Repository[T]) Get(id int) (*T, error) {
	var entity T
	err := r.DB.First(&entity, id).Error
	return &entity, err
}

func (r *Repository[T]) Update(entity *T) error {
	return r.DB.Save(entity).Error
}

func (r *Repository[T]) Delete(id int) error {
	return r.DB.Delete(new(T), id).Error
}