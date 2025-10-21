package subscription

import (
	"context"
	"subscription-service/internal/domain/models"
)

type SubscriptionStorage interface {
	Create(ctx context.Context, sub models.Subscription) (*models.Subscription, error)
	GetByID(ctx context.Context, id int64) (*models.Subscription, error)
	Update(ctx context.Context, sub models.Subscription) (*models.Subscription, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]models.Subscription, error)
	GetTotalPrice(ctx context.Context, userID, serviceName, startPeriod, endPeriod string) (int64, error)

	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type Service struct {
	db SubscriptionStorage
}

func NewService(db SubscriptionStorage) *Service {
	return &Service{
		db: db,
	}
}
