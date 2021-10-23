package database

import (
	"github.com/google/uuid"
	"github.com/sumelms/microservice-classroom/internal/classroomlesson/domain"
)

func toDBModel(entity *domain.ClassroomLesson) ClassroomLesson {
	cl := ClassroomLesson{
		ClassroomID: uuid.MustParse(entity.ClassroomID),
		LessonID:    uuid.MustParse(entity.LessonID),
	}

	if len(entity.UUID) > 0 {
		cl.UUID = uuid.MustParse(entity.UUID)
	}

	if entity.ID > 0 {
		// gorm.Model fields
		cl.ID = entity.ID
		cl.CreatedAt = entity.CreatedAt
		cl.UpdatedAt = entity.UpdatedAt

		if !entity.DeletedAt.IsZero() {
			cl.DeletedAt = entity.DeletedAt
		}
	}
	return cl
}

func toDomainModel(entity *ClassroomLesson) domain.ClassroomLesson {
	return domain.ClassroomLesson{
		ID:          entity.ID,
		UUID:        entity.UUID.String(),
		ClassroomID: entity.ClassroomID.String(),
		LessonID:    entity.LessonID.String(),
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
	}
}
