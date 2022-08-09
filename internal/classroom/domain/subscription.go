package domain

import "time"

type Subscription struct {
	ID          uint       `json:"id"`
	UUID        string     `json:"uuid"`
	UserID      string     `json:"user_id"`
	ClassroomID string     `json:"classroom_id"`
	Role        string     `json:"role"`
	ExpiresAt   *time.Time `json:"expires_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
