package subscriptionhand

import (
	"context"
	"net/http"
	"subscription-service/internal/domain/models"
	"subscription-service/internal/services/subscription"

	"github.com/rs/zerolog"
)

type Subscription interface {
	Create(ctx context.Context, sub models.Subscription) (*models.Subscription, error)
}

func Create(service *subscription.Service, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
