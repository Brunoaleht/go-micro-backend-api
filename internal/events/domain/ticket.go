package domain

import (
	"errors"

	"github.com/google/uuid"
)

type TicketStatus string

var (
	ErrTicketSpotRequired  = errors.New("ticket spot is required")
	ErrTicketPriceInvalid  = errors.New("ticket price must be greater than zero")
	ErrTicketStatusInvalid = errors.New("ticket status invalid")
)

const (
	TicketStatusHalf TicketStatus = "half"
	TicketStatusFull TicketStatus = "full"
)

func IsValidTicketStatus(status TicketStatus) bool {
	return status == TicketStatusHalf || status == TicketStatusFull
}

type Ticket struct {
	ID           string       `json:"id"`
	EventID      string       `json:"event_id"`
	Spot         *Spot        `json:"spot"`
	TicketKind TicketStatus `json:"ticket_status"`
	Price        float64      `json:"price"`
}


func CreatedNewTicket(e *Event, s *Spot, status TicketStatus) (*Ticket, error) {
	t := &Ticket{
		ID:           uuid.New().String(),
		EventID:      e.ID,
		Spot:         s,
		TicketKind: status,
		Price:        e.Price,
	}
	t.CalculatePrice()
	if err := t.Validate(); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Ticket) CalculatePrice() {
	if t.TicketKind == TicketStatusHalf {
		t.Price = t.Price / 2
	}
}

func (t Ticket) Validate() error {
	if t.Spot == nil {
		return ErrTicketSpotRequired
	}
	if t.Price <= 0 {
		return ErrTicketPriceInvalid
	}
	if !IsValidTicketStatus(t.TicketKind) {
		return ErrTicketStatusInvalid
	}
	return nil
}