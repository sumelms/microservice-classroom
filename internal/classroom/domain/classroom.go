package domain

import "time"

type Classroom struct {
	ID          uint
	UUID        string
	Title       string
	Subtitle    string
	Excerpt     string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
