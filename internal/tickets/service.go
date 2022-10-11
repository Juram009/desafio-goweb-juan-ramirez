package tickets

import (
	"context"
)

type Service interface {
	GetTotalTickets(ctx context.Context, destination string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	return len(tickets), nil
}

// COSTO PROMEDIO DE VALOR DE PASAJE AL DESTINO
func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0.0, err
	}
	totalPrice := 0.0
	for i := range tickets {
		totalPrice += tickets[i].Price
	}
	return totalPrice / float64(len(tickets)), nil
}
