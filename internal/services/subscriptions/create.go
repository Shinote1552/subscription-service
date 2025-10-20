package subscriptions

import (
	"context"
	"fmt"
	"subscription-service/internal/domain/models"
	"time"
)

func (s *Service) Create(ctx context.Context, sub models.Subscription) (*models.Subscription, error) {
	if sub.Price <= 0 {
		return nil, fmt.Errorf("invalid price: %d", sub.Price)
	}

	sub.UpdatedAt = time.Now()

	createdSub, err := s.db.Create(ctx, sub)
	if err != nil {
		return nil, fmt.Errorf("service create: %w", err)
	}

	return createdSub, nil
}
