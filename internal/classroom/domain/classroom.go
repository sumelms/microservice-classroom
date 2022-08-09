package domain

import (
	"time"

	"github.com/google/uuid"
)

type Classroom struct {
	ID          uint
	UUID        uuid.UUID
	Title       string
	Description string
	SubjectID   uuid.UUID
	CourseID    uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type ClassroomLesson struct {
	ID          uint       `json:"id"`
	UUID        uuid.UUID  `json:"uuid"`
	ClassroomID uuid.UUID  `json:"classroom_id"`
	LessonID    uuid.UUID  `json:"lesson_id"`
	StartsAt    time.Time  `json:"starts_at"`
	EndsAt      *time.Time `json:"ends_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
