package entity

import "time"

type Mother struct {
	MotherID       int        `gorm:"column:mother_id;primaryKey;autoIncrement"`
	UserID         int        `gorm:"column:user_id"`
	BloodType      *string    `gorm:"column:blood_type"`
	FullName       string     `gorm:"column:full_name"`
	BirthDate      *time.Time `gorm:"column:birth_date"`
	MedicalHistory *string    `gorm:"column:medical_history"`
	Height         *float64   `gorm:"column:height"`
}
