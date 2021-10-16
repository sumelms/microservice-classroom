package database

import (
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-classroom/internal/classroom/domain"
	merrors "github.com/sumelms/microservice-classroom/pkg/errors"
)

const (
	whereclassroomUUID = "uuid = ?"
)

// Repository struct
type Repository struct {
	db     *gorm.DB
	logger log.Logger
}

// NewRepository creates a new profile repository
func NewRepository(db *gorm.DB, logger log.Logger) *Repository {
	db.AutoMigrate(&Classroom{})

	return &Repository{
		db:     db,
		logger: logger,
	}
}

// List classrooms
func (r *Repository) List() ([]domain.Classroom, error) {
	var classrooms []Classroom

	query := r.db.Find(&classrooms)
	if query.RecordNotFound() {
		return []domain.Classroom{}, nil
	}
	if err := query.Error; err != nil {
		return []domain.Classroom{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "list classrooms")
	}

	var list []domain.Classroom
	for i := range classrooms {
		c := classrooms[i]
		list = append(list, toDomainModel(&c))
	}
	return list, nil
}

// Create creates a classroom
func (r *Repository) Create(classroom *domain.Classroom) (domain.Classroom, error) {
	entity := toDBModel(classroom)

	if err := r.db.Create(&entity).Error; err != nil {
		return domain.Classroom{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "can't create classroom")
	}
	return toDomainModel(&entity), nil
}

// Find get a classroom by its ID
func (r *Repository) Find(id string) (domain.Classroom, error) {
	var classroom Classroom

	query := r.db.Where(whereclassroomUUID, id).First(&classroom)
	if query.RecordNotFound() {
		return domain.Classroom{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "classroom not found")
	}
	if err := query.Error; err != nil {
		return domain.Classroom{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "find classroom")
	}

	return toDomainModel(&classroom), nil
}

// Update the given classroom
func (r *Repository) Update(c *domain.Classroom) (domain.Classroom, error) {
	var classroom Classroom

	query := r.db.Where(whereclassroomUUID, c.UUID).First(&classroom)

	if query.RecordNotFound() {
		return domain.Classroom{}, merrors.NewErrorf(merrors.ErrCodeNotFound, "classroom not found")
	}

	query = r.db.Model(&classroom).Update(&c)

	if err := query.Error; err != nil {
		return domain.Classroom{}, merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "can't update classroom")
	}

	return *c, nil
}

// Delete a classroom by its ID
func (r *Repository) Delete(id string) error {
	query := r.db.Where(whereclassroomUUID, id).Delete(&Classroom{})

	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return merrors.WrapErrorf(err, merrors.ErrCodeNotFound, "classroom not found")
		}
		return merrors.WrapErrorf(err, merrors.ErrCodeUnknown, "delete classroom")
	}

	return nil
}
