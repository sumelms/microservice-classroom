package domain

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
)

type ServiceInterface interface {
	ListClassroom(context.Context) ([]Classroom, error)
	CreateClassroom(context.Context, *Classroom) (Classroom, error)
	FindClassroom(context.Context, string) (Classroom, error)
	UpdateClassroom(context.Context, *Classroom) (Classroom, error)
	DeleteClassroom(context.Context, string) error
}

type Service struct {
	repo   RepositoryInterface
	logger log.Logger
}

func NewService(repo RepositoryInterface, logger log.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

func (s *Service) ListClassroom(_ context.Context) ([]Classroom, error) {
	cs, err := s.repo.List()
	if err != nil {
		return []Classroom{}, fmt.Errorf("Service didn't found any classroom: %w", err)
	}
	return cs, nil
}

func (s *Service) CreateClassroom(_ context.Context, classroom *Classroom) (Classroom, error) {
	c, err := s.repo.Create(classroom)
	if err != nil {
		return Classroom{}, fmt.Errorf("Service can't create classroom: %w", err)
	}
	return c, nil
}

func (s *Service) FindClassroom(_ context.Context, id string) (Classroom, error) {
	c, err := s.repo.Find(id)
	if err != nil {
		return Classroom{}, fmt.Errorf("Service can't find classroom: %w", err)
	}
	return c, nil
}

func (s *Service) UpdateClassroom(_ context.Context, classroom *Classroom) (Classroom, error) {
	c, err := s.repo.Update(classroom)
	if err != nil {
		return Classroom{}, fmt.Errorf("Service can't update classroom: %w", err)
	}
	return c, nil
}

func (s *Service) DeleteClassroom(_ context.Context, id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("Service can't delete classroom: %w", err)
	}
	return nil
}
