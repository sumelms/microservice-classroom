package domain

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
)

type ServiceInterface interface {
	ListClassroomLesson(context.Context) ([]ClassroomLesson, error)
	CreateClassroomLesson(context.Context, *ClassroomLesson) (ClassroomLesson, error)
	FindClassroomLesson(context.Context, string) (ClassroomLesson, error)
	UpdateClassroomLesson(context.Context, *ClassroomLesson) (ClassroomLesson, error)
	DeleteClassroomLesson(context.Context, string) error
}

type Service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

func (s *Service) ListClassroomLesson(_ context.Context) ([]ClassroomLesson, error) {
	cs, err := s.repo.List()
	if err != nil {
		return []ClassroomLesson{}, fmt.Errorf("Service didn't found any classroom: %w", err)
	}
	return cs, nil
}

func (s *Service) CreateClassroomLesson(_ context.Context, classroom *ClassroomLesson) (ClassroomLesson, error) {
	c, err := s.repo.Create(classroom)
	if err != nil {
		return ClassroomLesson{}, fmt.Errorf("Service can't create classroom: %w", err)
	}
	return c, nil
}

func (s *Service) FindClassroomLesson(_ context.Context, id string) (ClassroomLesson, error) {
	c, err := s.repo.Find(id)
	if err != nil {
		return ClassroomLesson{}, fmt.Errorf("Service can't find classroom: %w", err)
	}
	return c, nil
}

func (s *Service) UpdateClassroomLesson(_ context.Context, classroom *ClassroomLesson) (ClassroomLesson, error) {
	c, err := s.repo.Update(classroom)
	if err != nil {
		return ClassroomLesson{}, fmt.Errorf("Service can't update classroom: %w", err)
	}
	return c, nil
}

func (s *Service) DeleteClassroomLesson(_ context.Context, id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("Service can't delete classroom: %w", err)
	}
	return nil
}
