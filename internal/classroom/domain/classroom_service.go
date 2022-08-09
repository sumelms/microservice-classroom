package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) Classroom(_ context.Context, id uuid.UUID) (Classroom, error) {
	c, err := s.classrooms.Classroom(id)
	if err != nil {
		return Classroom{}, fmt.Errorf("Service can't find classroom: %w", err)
	}
	return c, nil
}

func (s *Service) Classrooms(_ context.Context) ([]Classroom, error) {
	cs, err := s.classrooms.Classrooms()
	if err != nil {
		return []Classroom{}, fmt.Errorf("Service didn't found any classroom: %w", err)
	}
	return cs, nil
}

func (s *Service) CreateClassroom(_ context.Context, classroom *Classroom) (Classroom, error) {
	c, err := s.classrooms.CreateClassroom(classroom)
	if err != nil {
		return Classroom{}, fmt.Errorf("Service can't create classroom: %w", err)
	}
	return c, nil
}

func (s *Service) UpdateClassroom(_ context.Context, classroom *Classroom) (Classroom, error) {
	c, err := s.classrooms.UpdateClassroom(classroom)
	if err != nil {
		return Classroom{}, fmt.Errorf("Service can't update classroom: %w", err)
	}
	return c, nil
}

func (s *Service) DeleteClassroom(_ context.Context, id uuid.UUID) error {
	err := s.classrooms.DeleteClassroom(id)
	if err != nil {
		return fmt.Errorf("Service can't delete classroom: %w", err)
	}
	return nil
}

func (s *Service) AddLesson(ctx context.Context, classroomID, lessonID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) RemoveLesson(ctx context.Context, classroomID, lessonID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
