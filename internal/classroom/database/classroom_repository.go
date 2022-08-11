package database

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-classroom/internal/classroom/domain"
	"github.com/sumelms/microservice-classroom/pkg/errors"
)

// NewClassroomRepository creates the subject subjectRepository
func NewClassroomRepository(db *sqlx.DB) (classroomRepository, error) { //nolint: revive
	sqlStatements := make(map[string]*sqlx.Stmt)

	for queryName, query := range queriesClassroom() {
		stmt, err := db.Preparex(query)
		if err != nil {
			return classroomRepository{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement %s", queryName)
		}
		sqlStatements[queryName] = stmt
	}

	return classroomRepository{
		statements: sqlStatements,
	}, nil
}

type classroomRepository struct {
	statements map[string]*sqlx.Stmt
}

func (c classroomRepository) Classroom(id uuid.UUID) (domain.Classroom, error) {
	//TODO implement me
	panic("implement me")
}

func (c classroomRepository) Classrooms() ([]domain.Classroom, error) {
	//TODO implement me
	panic("implement me")
}

func (c classroomRepository) CreateClassroom(classroom *domain.Classroom) (domain.Classroom, error) {
	//TODO implement me
	panic("implement me")
}

func (c classroomRepository) UpdateClassroom(classroom *domain.Classroom) (domain.Classroom, error) {
	//TODO implement me
	panic("implement me")
}

func (c classroomRepository) DeleteClassroom(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
