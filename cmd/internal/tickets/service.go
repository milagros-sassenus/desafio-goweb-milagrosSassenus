package tickets

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetTotalTickets(c *gin.Context, destination string) (int, error)
	AverageDestination(c *gin.Context, destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

func (s *service) GetTotalTickets(ctx *gin.Context, destination string) (int, error) {
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	return len(tickets), nil
}

func (s *service) AverageDestination(ctx *gin.Context, destination string) (float64, error) {

	totalTickets, errTotal := s.repository.GetAll(ctx)
	if errTotal != nil {
		return 0, errTotal
	}

	tickets, errTicket := s.repository.GetTicketByDestination(ctx, destination)
	if errTicket != nil {
		return 0, errTicket
	}

	return float64(len(tickets)) / float64(len(totalTickets)), nil
}
