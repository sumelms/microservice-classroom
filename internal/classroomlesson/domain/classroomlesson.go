package domain

import "time"

type ClassroomLesson struct {
	ID          uint       `json:"id"`
	UUID        string     `json:"uuid"`
	ClassroomID string     `json:"classroom_id"`
	LessonID    string     `json:"lesson_id"`
	StartsAt    time.Time  `json:"starts_at"`
	EndsAt      *time.Time `json:"ends_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
