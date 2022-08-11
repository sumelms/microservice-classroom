package domain

import (
	"context"

	"github.com/go-kit/log"
	"github.com/google/uuid"
)

type ServiceInterface interface {
	Classroom(ctx context.Context, id uuid.UUID) (Classroom, error)
	Classrooms(ctx context.Context) ([]Classroom, error)
	CreateClassroom(ctx context.Context, classroom *Classroom) (Classroom, error)
	UpdateClassroom(ctx context.Context, classroom *Classroom) (Classroom, error)
	DeleteClassroom(ctx context.Context, id uuid.UUID) error
	AddLesson(ctx context.Context, classroomID, lessonID uuid.UUID) error
	RemoveLesson(ctx context.Context, classroomID, lessonID uuid.UUID) error

	Subscription(ctx context.Context, id uuid.UUID) (Subscription, error)
	Subscriptions(ctx context.Context) ([]Subscription, error)
	CreateSubscription(ctx context.Context, subscription *Subscription) (Subscription, error)
	UpdateSubscription(ctx context.Context, subscription *Subscription) (Subscription, error)
	DeleteSubscription(ctx context.Context, id uuid.UUID) error
}

type serviceConfiguration func(svc *Service) error

type Service struct {
	classrooms    ClassroomRepository
	subscriptions SubscriptionRepository
	logger        log.Logger
}

func NewService(cfgs ...serviceConfiguration) (*Service, error) {
	svc := &Service{}
	for _, cfg := range cfgs {
		err := cfg(svc)
		if err != nil {
			return nil, err
		}
	}
	return svc, nil
}

// WithClassroomRepository injects the course repository to the domain Service
func WithClassroomRepository(cr ClassroomRepository) serviceConfiguration {
	return func(svc *Service) error {
		svc.classrooms = cr
		return nil
	}
}

// WithSubscriptionRepository injects the course repository to the domain Service
func WithSubscriptionRepository(sr SubscriptionRepository) serviceConfiguration {
	return func(svc *Service) error {
		svc.subscriptions = sr
		return nil
	}
}

// WithLogger injects the logger to the domain Service
func WithLogger(l log.Logger) serviceConfiguration {
	return func(svc *Service) error {
		svc.logger = l
		return nil
	}
}
