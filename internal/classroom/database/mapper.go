package database

import (
	"github.com/google/uuid"
	"github.com/sumelms/microservice-classroom/internal/classroom/domain"
)

func toDBModel(entity *domain.Classroom) Classroom {
	classroom := Classroom{
		Title:       entity.Title,
		Subtitle:    entity.Subtitle,
		Excerpt:     entity.Excerpt,
		Description: entity.Description,
	}

	if len(entity.UUID) > 0 {
		classroom.UUID = uuid.MustParse(entity.UUID)
	}

	if entity.ID > 0 {
		// gorm.Model fields
		classroom.ID = entity.ID
		classroom.CreatedAt = entity.CreatedAt
		classroom.UpdatedAt = entity.UpdatedAt

		if !entity.DeletedAt.IsZero() {
			classroom.DeletedAt = entity.DeletedAt
		}
	}
	return classroom
}

func toDomainModel(entity *Classroom) domain.Classroom {
	return domain.Classroom{
		ID:          entity.ID,
		UUID:        entity.UUID.String(),
		Title:       entity.Title,
		Subtitle:    entity.Subtitle,
		Excerpt:     entity.Excerpt,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
	}
}
