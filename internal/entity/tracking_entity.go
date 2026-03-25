package entity

import "time"

type Question struct {
	QuestionID   int       `gorm:"column:question_id;primaryKey;autoIncrement"`
	ScheduleID   *int      `gorm:"column:schedule_id"`
	QuestionText string    `gorm:"column:question_text"`
	QuestionKey  string    `gorm:"column:question_key"`
	Category     *string   `gorm:"column:category"`
	InputType    *string   `gorm:"column:input_type"`
	Unit         *string   `gorm:"column:unit"`
	IsRequired   *bool     `gorm:"column:is_required;default:true"`
	DisplayOrder *int      `gorm:"column:display_order"`
	CreatedAt    time.Time `gorm:"column:created_at"`
}

type QuestionOption struct {
	OptionID     int     `gorm:"column:option_id;primaryKey;autoIncrement"`
	QuestionID   *int    `gorm:"column:question_id"`
	OptionValue  string  `gorm:"column:option_value"`
	OptionLabel  string  `gorm:"column:option_label"`
	IconEmoji    *string `gorm:"column:icon_emoji"`
	DisplayOrder *int    `gorm:"column:display_order"`
	IsDefault    *bool   `gorm:"column:is_default;default:false"`
}

type UserTrackingLog struct {
	TrackingID       int       `gorm:"column:tracking_id;primaryKey;autoIncrement"`
	MotherID         *int      `gorm:"column:mother_id"`
	PregnancyID      *int      `gorm:"column:pregnancy_id"`
	ChildID          *int      `gorm:"column:child_id"`
	TrackingDate     time.Time `gorm:"column:tracking_date"`
	Frequency        string    `gorm:"column:frequency"`
	
	AnswersData      string    `gorm:"column:answers_data;type:json"`
	
	CompletionStatus *string   `gorm:"column:completion_status"`
	Notes            *string   `gorm:"column:notes"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
}
