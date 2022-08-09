package domain

import "github.com/google/uuid"

type ClassroomRepository interface {
	Classroom(id uuid.UUID) (Classroom, error)
	Classrooms() ([]Classroom, error)
	CreateClassroom(classroom *Classroom) (Classroom, error)
	UpdateClassroom(classroom *Classroom) (Classroom, error)
	DeleteClassroom(id uuid.UUID) error
}
