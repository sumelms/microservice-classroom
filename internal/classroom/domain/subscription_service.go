package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) ListSubscription(_ context.Context) ([]Subscription, error) {
	list, err := s.subscriptions.Subscriptions()
	if err != nil {
		return []Subscription{}, fmt.Errorf("Service didn't found any subscription: %w", err)
	}
	return list, nil
}

func (s *Service) CreateSubscription(_ context.Context, subscription *Subscription) (Subscription, error) {
	sub, err := s.subscriptions.CreateSubscription(subscription)
	if err != nil {
		return Subscription{}, fmt.Errorf("Service can't create subscription: %w", err)
	}
	return sub, nil
}

func (s *Service) FindSubscription(_ context.Context, id uuid.UUID) (Subscription, error) {
	sub, err := s.subscriptions.Subscription(id)
	if err != nil {
		return Subscription{}, fmt.Errorf("Service can't find subscription: %w", err)
	}
	return sub, nil
}

func (s *Service) UpdateSubscription(_ context.Context, subscription *Subscription) (Subscription, error) {
	sub, err := s.subscriptions.UpdateSubscription(subscription)
	if err != nil {
		return Subscription{}, fmt.Errorf("Service can't update subscription: %w", err)
	}
	return sub, nil
}

func (s *Service) DeleteSubscription(ctx context.Context, id uuid.UUID) error {
	err := s.subscriptions.DeleteSubscription(id)
	if err != nil {
		return fmt.Errorf("Service can't delete subscription: %w", err)
	}
	return nil
}
