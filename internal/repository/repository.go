package repository

import gorm "gorm.io/gorm"

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *Repository[T]) List(db *gorm.DB) ([]T, error) {
	var entities []T
	err := db.Find(&entities).Error
	return entities, err
	
}

func (r *Repository[T]) Get(db *gorm.DB, id int) (*T, error) {
	var entity T
	err := db.First(&entity, id).Error
	return &entity, err
}

func (r *Repository[T]) Update(db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, id int) error {
	return db.Delete(new(T), id).Error
}