package domain

import (
	"time"

	"github.com/google/uuid"
)

type Classroom struct {
	ID          uint       `json:"id"`
	UUID        uuid.UUID  `json:"uuid"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	Format      string     `json:"format"`
	SubjectID   *uuid.UUID `db:"subject_id" json:"subject_id,omitempty"`
	CourseID    uuid.UUID  `db:"course_id" json:"course_id"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

type ClassroomLesson struct {
	ID          uint       `json:"id"`
	UUID        uuid.UUID  `json:"uuid"`
	ClassroomID uuid.UUID  `db:"classroom_id" json:"classroom_id"`
	LessonID    uuid.UUID  `db:"lesson_id" json:"lesson_id"`
	StartsAt    time.Time  `db:"starts_at" json:"starts_at"`
	EndsAt      *time.Time `db:"ends_at" json:"ends_at,omitempty"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}
