package subscriptions

import (
	"context"
	"subscription-service/internal/domain/models"
)

type SubscriptionRepository interface {
	Create(ctx context.Context, sub models.Subscription) (*models.Subscription, error)
	GetByID(ctx context.Context, id int64) (*models.Subscription, error)
	Update(ctx context.Context, sub models.Subscription) (*models.Subscription, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]models.Subscription, error)
	GetTotalPrice(ctx context.Context, userID, serviceName, startPeriod, endPeriod string) (int64, error)
}

type Service struct {
	db SubscriptionRepository
}

func NewService(db SubscriptionRepository) *Service {
	return &Service{
		db: db,
	}
}
