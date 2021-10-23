package database

import (
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-classroom/internal/classroomlesson/domain"
	merrors "github.com/sumelms/microservice-classroom/pkg/errors"
)

const (
	whereClassroomLessonUUID = "uuid = ?"
)

// Repository struct
type Repository struct {
	db     *gorm.DB
	logger log.Logger
}

// NewRepository creates a new profile repository
func NewRepository(db *gorm.DB, logger log.Logger) *Repository {
	db.AutoMigrate(&ClassroomLesson{})

	return &Repository{
		db:     db,
		logger: logger,
	}
}

// List classroomlessons
func (r *Repository) List() ([]domain.ClassroomLesson, error) {
	var classrooms []ClassroomLesson

	query := r.db.Find(&classrooms)
	if query.RecordNotFound() {
		return []domain.ClassroomLesson{}, nil
	}
	if err := query.Error; err != nil {
		return []domain.ClassroomLesson{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "list classroomlessons")
	}

	var list []domain.ClassroomLesson
	for i := range classrooms {
		c := classrooms[i]
		list = append(list, toDomainModel(&c))
	}
	return list, nil
}

// Create creates a classroomlesson
func (r *Repository) Create(classroom *domain.ClassroomLesson) (domain.ClassroomLesson, error) {
	entity := toDBModel(classroom)

	if err := r.db.Create(&entity).Error; err != nil {
		return domain.ClassroomLesson{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "can't create classroomlesson")
	}
	return toDomainModel(&entity), nil
}

// Find get a classroomlesson by its ID
func (r *Repository) Find(id string) (domain.ClassroomLesson, error) {
	var classroom ClassroomLesson

	query := r.db.Where(whereClassroomLessonUUID, id).First(&classroom)
	if query.RecordNotFound() {
		return domain.ClassroomLesson{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "classroom not found")
	}
	if err := query.Error; err != nil {
		return domain.ClassroomLesson{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "find classroomlesson")
	}

	return toDomainModel(&classroom), nil
}

// Update the given classroomlesson
func (r *Repository) Update(c *domain.ClassroomLesson) (domain.ClassroomLesson, error) {
	var classroom ClassroomLesson

	query := r.db.Where(whereClassroomLessonUUID, c.UUID).First(&classroom)

	if query.RecordNotFound() {
		return domain.ClassroomLesson{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "classroomlesson not found")
	}

	query = r.db.Model(&classroom).Update(&c)

	if err := query.Error; err != nil {
		return domain.ClassroomLesson{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "can't update classroomlesson")
	}

	return *c, nil
}

// Delete a classroom by its ID
func (r *Repository) Delete(id string) error {
	query := r.db.Where(whereClassroomLessonUUID, id).Delete(&ClassroomLesson{})

	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return merrors.WrapErrorf(err, merrors.ErrCodeNotFound, "classroomlesson not found")
		}
		return merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "delete classroomlesson")
	}

	return nil
}
