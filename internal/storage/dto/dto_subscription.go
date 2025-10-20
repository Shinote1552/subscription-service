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

func ToDomain(dto SubscriptionDB) models.Subscription {
	return models.Subscription{
		ID:          dto.ID,
		ServiceName: dto.ServiceName,
		Price:       dto.Price,
		UserID:      dto.UserID,
		StartDate:   dto.StartDate,
		EndDate:     dto.EndDate,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
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
