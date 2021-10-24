package domain

import "time"

type Lesson struct {
	ID          uint
	UUID        string
	Title       string
	Subtitle    string
	Excerpt     string
	Description string
	Module      string
	SubjectID   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
