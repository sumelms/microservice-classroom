package domain

import "time"

type ClassroomLesson struct {
	ID          uint
	UUID        string
	ClassroomID string
	LessonID    string
	StartsAt    time.Time
	EndsAt      *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
