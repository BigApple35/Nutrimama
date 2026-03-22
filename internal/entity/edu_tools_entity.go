package entity

import "time"

type EduTools struct {
	ID          int    `gorm:"column:edu_tools_id;primaryKey;autoIncrement"`
	Publisher   string `gorm:"column:publisher"`
	Title      string `gorm:"column:title"`
	Content        string `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:created_at"`
}