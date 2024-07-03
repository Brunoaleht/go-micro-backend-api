package usecase

import (
	"go-backend-api/internal/events/domain"
)

type BuyTicketsInputDto struct {
	EventID string `json:"event_id"`
	Spots []string `json:"spots"`
	TicketKind string `json:"ticket_kind"`
	CardHash string `json:"card_hash"`
	Email string `json:"email"`
}

type BuyTicketsOutputDto struct {
	Tickets []TicketDto `json:"tickets"`
}

type TicketDto struct {
	ID string `json:"id"`
	SpotID string `json:"spot_id"`
	EventID string `json:"event_id"`
	TicketKind string `json:"ticket_kind"`
	Price float64 `json:"price"`
}

type BuyTicketsUseCase struct {
	repo domain.EventRepository
}