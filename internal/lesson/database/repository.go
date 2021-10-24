package database

import (
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-classroom/internal/lesson/domain"
	merrors "github.com/sumelms/microservice-classroom/pkg/errors"
)

const (
	whereLessonUUID = "uuid = ?"
)

// Repository struct
type Repository struct {
	db     *gorm.DB
	logger log.Logger
}

// NewRepository creates a new lesson repository
func NewRepository(db *gorm.DB, logger log.Logger) *Repository {
	db.AutoMigrate(&Lesson{})

	return &Repository{
		db:     db,
		logger: logger,
	}
}

// List lessons
func (r *Repository) List(filters map[string]interface{}) ([]domain.Lesson, error) {
	var lessons []Lesson

	query := r.db.Find(&lessons, filters)
	if query.RecordNotFound() {
		return []domain.Lesson{}, nil
	}
	if err := query.Error; err != nil {
		return []domain.Lesson{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "list lessons")
	}

	var list []domain.Lesson
	for i := range lessons {
		l := lessons[i]
		list = append(list, toDomainModel(&l))
	}
	return list, nil
}

// Create creates a lesson
func (r *Repository) Create(lesson *domain.Lesson) (domain.Lesson, error) {
	entity := toDBModel(lesson)

	if err := r.db.Create(&entity).Error; err != nil {
		return domain.Lesson{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "can't create lesson")
	}
	return toDomainModel(&entity), nil
}

// Find get a lesson by its ID
func (r *Repository) Find(id string) (domain.Lesson, error) {
	var lesson Lesson

	query := r.db.Where(whereLessonUUID, id).First(&lesson)
	if query.RecordNotFound() {
		return domain.Lesson{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "lesson not found")
	}
	if err := query.Error; err != nil {
		return domain.Lesson{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "find lesson")
	}

	return toDomainModel(&lesson), nil
}

// Update the given lesson
func (r *Repository) Update(c *domain.Lesson) (domain.Lesson, error) {
	var lesson Lesson

	query := r.db.Where(whereLessonUUID, c.UUID).First(&lesson)

	if query.RecordNotFound() {
		return domain.Lesson{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "lesson not found")
	}

	query = r.db.Model(&lesson).Update(&c)

	if err := query.Error; err != nil {
		return domain.Lesson{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "can't update lesson")
	}

	return *c, nil
}

// Delete a lesson by its ID
func (r *Repository) Delete(id string) error {
	query := r.db.Where(whereLessonUUID, id).Delete(&Lesson{})

	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return merrors.WrapErrorf(err, merrors.ErrCodeNotFound, "lesson not found")
		}
		return merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "delete lesson")
	}

	return nil
}
