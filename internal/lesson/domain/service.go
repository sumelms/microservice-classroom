package domain

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
)

type ServiceInterface interface {
	ListLesson(context.Context) ([]Lesson, error)
	CreateLesson(context.Context, *Lesson) (Lesson, error)
	FindLesson(context.Context, string) (Lesson, error)
	UpdateLesson(context.Context, *Lesson) (Lesson, error)
	DeleteLesson(context.Context, string) error
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

func (s *Service) ListLesson(_ context.Context) ([]Lesson, error) {
	cs, err := s.repo.List()
	if err != nil {
		return []Lesson{}, fmt.Errorf("Service didn't found any lesson: %w", err)
	}
	return cs, nil
}

func (s *Service) CreateLesson(_ context.Context, classroom *Lesson) (Lesson, error) {
	l, err := s.repo.Create(classroom)
	if err != nil {
		return Lesson{}, fmt.Errorf("Service can't create lesson: %w", err)
	}
	return l, nil
}

func (s *Service) FindLesson(_ context.Context, id string) (Lesson, error) {
	l, err := s.repo.Find(id)
	if err != nil {
		return Lesson{}, fmt.Errorf("Service can't find lesson: %w", err)
	}
	return l, nil
}

func (s *Service) UpdateLesson(_ context.Context, classroom *Lesson) (Lesson, error) {
	l, err := s.repo.Update(classroom)
	if err != nil {
		return Lesson{}, fmt.Errorf("Service can't update lesson: %w", err)
	}
	return l, nil
}

func (s *Service) DeleteLesson(_ context.Context, id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("Service can't delete lesson: %w", err)
	}
	return nil
}
