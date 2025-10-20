package get

import (
	"subscription-service/internal/domain/models"
	"time"

	"github.com/google/uuid"
)

type Response struct {
	ID          int64     `json:"id"`
	ServiceName string    `json:"service_name"`
	Price       int64     `json:"price"`
	UserID      uuid.UUID `json:"user_id"`
	StartDate   string    `json:"start_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(domain models.Subscription) Response {
	return Response{
		ID:          domain.ID,
		ServiceName: domain.ServiceName,
		Price:       domain.Price,
		UserID:      domain.UserID,
		StartDate:   domain.StartDate.Format("01-2006"),
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
