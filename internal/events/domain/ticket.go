package domain

type TicketStatus string

const (
	TicketStatusHalf TicketStatus = "half"
	TicketStatusFull TicketStatus = "full"
)

type Ticket struct {
	ID           string       `json:"id"`
	EventID      string       `json:"event_id"`
	Spot         *Spot        `json:"spot"`
	TicketStatus TicketStatus `json:"ticket_status"`
	Price        float64      `json:"price"`
}