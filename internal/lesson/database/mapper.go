package database

import (
	"github.com/google/uuid"
	"github.com/sumelms/microservice-classroom/internal/lesson/domain"
)

func toDBModel(entity *domain.Lesson) Lesson {
	lesson := Lesson{
		Title:       entity.Title,
		Subtitle:    entity.Subtitle,
		Excerpt:     entity.Excerpt,
		Description: entity.Description,
		Module:      entity.Module,
		SubjectID:   uuid.MustParse(entity.SubjectID),
	}

	if len(entity.UUID) > 0 {
		lesson.UUID = uuid.MustParse(entity.UUID)
	}

	if entity.ID > 0 {
		// gorm.Model fields
		lesson.ID = entity.ID
		lesson.CreatedAt = entity.CreatedAt
		lesson.UpdatedAt = entity.UpdatedAt

		if !entity.DeletedAt.IsZero() {
			lesson.DeletedAt = entity.DeletedAt
		}
	}
	return lesson
}

func toDomainModel(entity *Lesson) domain.Lesson {
	return domain.Lesson{
		ID:          entity.ID,
		UUID:        entity.UUID.String(),
		Title:       entity.Title,
		Subtitle:    entity.Subtitle,
		Excerpt:     entity.Excerpt,
		Description: entity.Description,
		Module:      entity.Module,
		SubjectID:   entity.SubjectID.String(),
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
	}
}
