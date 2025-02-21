package models

import "time"

type StudySession struct {
	ID              int       `json:"id"`
	GroupID         int       `json:"group_id"`
	CreatedAt       time.Time `json:"created_at"`
	StudyActivityID int       `json:"study_activity_id"`
	ActivityName    string    `json:"activity_name"`
	GroupName       string    `json:"group_name"`
}
