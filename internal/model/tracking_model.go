package model

import "time"

type SubmitTrackingRequest struct {
	Frequency    string    `json:"frequency" binding:"required"` // daily, weekly, monthly
	TrackingDate time.Time `json:"tracking_date" binding:"required"`
	AnswersData  string    `json:"answers_data" binding:"required"` // The JSON payload of chosen keys
	Notes        *string   `json:"notes"`
}
