package domain

import "github.com/google/uuid"

type SubscriptionRepository interface {
	Subscription(id uuid.UUID) (Subscription, error)
	Subscriptions() ([]Subscription, error)
	CreateSubscription(subscription *Subscription) (Subscription, error)
	UpdateSubscription(subscription *Subscription) (Subscription, error)
	DeleteSubscription(id uuid.UUID) error
}
