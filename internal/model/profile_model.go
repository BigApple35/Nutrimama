package model

import "time"

type EditMotherProfileRequest struct {
	FullName       string     `json:"full_name" binding:"required"`
	BloodType      *string    `json:"blood_type"`
	BirthDate      *time.Time `json:"birth_date"`
	MedicalHistory *string    `json:"medical_history"`
	Height         *float64   `json:"height"`
}

type EditConsultantProfileRequest struct {
	FullName     string  `json:"full_name" binding:"required"`
	FacilityName *string `json:"facility_name"`
}
