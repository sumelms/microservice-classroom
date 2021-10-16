package domain

import "time"

type Lesson struct {
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
