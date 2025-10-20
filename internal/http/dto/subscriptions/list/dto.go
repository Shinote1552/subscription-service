package list

import (
	"subscription-service/internal/domain/models"
	"subscription-service/internal/http/dto/subscriptions/get"
)

type Response struct {
	Subscriptions []get.Response `json:"subscriptions"`
}

func FromDomain(subscriptions []models.Subscription) Response {
	responses := make([]get.Response, len(subscriptions))
	for i, sub := range subscriptions {
		responses[i] = get.FromDomain(sub)
	}
	return Response{Subscriptions: responses}
}
