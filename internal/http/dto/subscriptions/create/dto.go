package create

import (
	"fmt"
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

func ToDomain(subDTO Request) (models.Subscription, error) {
	startDateStringed, err := time.Parse("01-2006`", subDTO.StartDate)
	if err != nil {
		return models.Subscription{}, fmt.Errorf("to domain converting: %w", err)
	}

	return models.Subscription{
			ServiceName: subDTO.ServiceName,
			Price:       subDTO.Price,
			UserID:      subDTO.UserID,
			StartDate:   startDateStringed,
		},
		nil
}

func FromDomain(subDomain models.Subscription) Request {
	startDate := subDomain.StartDate.Format(time.RFC3339)

	return Request{
		ServiceName: subDomain.ServiceName,
		Price:       subDomain.Price,
		UserID:      subDomain.UserID,
		StartDate:   startDate,
	}

}
