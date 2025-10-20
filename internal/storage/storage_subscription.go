package storage

import (
	"context"
	"subscription-service/internal/domain/models"
	"subscription-service/internal/storage/dto"
)

func (s *Storage) Create(ctx context.Context, sub models.Subscription) (*models.Subscription, error) {
	dtoSub := dto.FromDomain(sub)

	query := `
		INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	err := s.pool.QueryRow(ctx, query,
		dtoSub.ServiceName, dtoSub.Price, dtoSub.UserID, dtoSub.StartDate, dtoSub.EndDate,
	).Scan(&dtoSub.ID, &dtoSub.CreatedAt, &dtoSub.UpdatedAt)

	if err != nil {
		return nil, err
	}

	modelSub := dto.ToDomain(dtoSub)
	return &modelSub, nil
}

func (s *Storage) GetByID(ctx context.Context, id int64) (*models.Subscription, error)
func (s *Storage) Update(ctx context.Context, sub models.Subscription) (*models.Subscription, error)
func (s *Storage) Delete(ctx context.Context, id int64) error
func (s *Storage) List(ctx context.Context) ([]models.Subscription, error)
func (s *Storage) GetTotalPrice(ctx context.Context, userID, serviceName, startPeriod, endPeriod string) (int64, error)
