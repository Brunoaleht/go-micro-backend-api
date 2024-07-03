package domain

type EventRepository interface {
	ListEvents() ([]Event, error)
	GetEventByID(eventId string) (*Event, error)
	FindSpotsEventID(eventId string) ([]Spot, error)
	FindSpotByNames(eventId, spotNames string) ([]Spot, error)
	CreateSpot(spot *Spot) error
	CreateTicket(ticket *Ticket) error
	ReserveSpot(spotId, ticketId string) error
	CreateEvent(event *Event) error
}