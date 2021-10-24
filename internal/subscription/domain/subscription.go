package domain

import "time"

type Subscription struct {
	ID          uint
	UUID        string
	UserID      string
	ClassroomID string
	Role        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
