package database

import (
	"github.com/google/uuid"
	"github.com/sumelms/microservice-classroom/internal/subscription/domain"
)

func toDBModel(entity *domain.Subscription) Subscription {
	s := Subscription{
		UserID:      uuid.MustParse(entity.UserID),
		ClassroomID: uuid.MustParse(entity.ClassroomID),
		Role:        entity.Role,
	}

	if len(entity.UUID) > 0 {
		s.UUID = uuid.MustParse(entity.UUID)
	}

	if entity.ID > 0 {
		// gorm.Model fields
		s.ID = entity.ID
		s.CreatedAt = entity.CreatedAt
		s.UpdatedAt = entity.UpdatedAt

		if !entity.DeletedAt.IsZero() {
			s.DeletedAt = entity.DeletedAt
		}
	}

	return s
}

func toDomainModel(entity *Subscription) domain.Subscription {
	return domain.Subscription{
		ID:          entity.ID,
		UUID:        entity.UUID.String(),
		UserID:      entity.UserID.String(),
		ClassroomID: entity.ClassroomID.String(),
		Role:        entity.Role,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
	}
}
