package postgres

import (
	"context"
	"subscription-service/internal/domain/models"
)

func (p *PostgresStorage) Create(ctx context.Context, sub models.Subscription) (*models.Subscription, error) {

}
