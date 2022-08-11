package domain

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          uint       `json:"id"`
	UUID        uuid.UUID  `json:"uuid"`
	UserID      uuid.UUID  `json:"user_id"`
	ClassroomID uuid.UUID  `json:"classroom_id"`
	Role        string     `json:"role"`
	ExpiresAt   *time.Time `json:"expires_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
