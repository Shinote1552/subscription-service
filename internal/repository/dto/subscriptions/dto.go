package dto

import (
	"subscription-service/internal/domain/models"
	"time"

	"github.com/google/uuid"
)

type SubscriptionDB struct {
	ID          int64
	ServiceName string
	Price       int64
	UserID      uuid.UUID
	StartDate   time.Time
	EndDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (d SubscriptionDB) ToDomain() models.Subscription {
	return models.Subscription{
		ID:          d.ID,
		ServiceName: d.ServiceName,
		Price:       d.Price,
		UserID:      d.UserID,
		StartDate:   d.StartDate,
		EndDate:     d.EndDate,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

func FromDomain(domain models.Subscription) SubscriptionDB {
	return SubscriptionDB{
		ID:          domain.ID,
		ServiceName: domain.ServiceName,
		Price:       domain.Price,
		UserID:      domain.UserID,
		StartDate:   domain.StartDate,
		EndDate:     domain.EndDate,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
