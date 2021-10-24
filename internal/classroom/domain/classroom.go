package domain

import (
	"time"
)

type Classroom struct {
	ID          uint
	UUID        string
	Title       string
	Description string
	SubjectID   string
	CourseID    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
