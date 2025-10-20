package create

import (
	"subscription-service/internal/domain/models"
	"time"

	"github.com/google/uuid"
)

type Request struct {
	ServiceName string    `json:"service_name"`
	Price       int64     `json:"price"`
	UserID      uuid.UUID `json:"user_id"`
	StartDate   string    `json:"start_date"`
}

func (r Request) ToDomain() models.Subscription {
	startDate, _ := time.Parse("01-2006", r.StartDate)

	return models.Subscription{
		ServiceName: r.ServiceName,
		Price:       r.Price,
		UserID:      r.UserID,
		StartDate:   startDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
