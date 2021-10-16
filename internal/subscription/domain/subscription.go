package domain

import "time"

type Subscription struct {
	ID         uint
	UserID     string
	ClassroomID   string
	ValidUntil *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
