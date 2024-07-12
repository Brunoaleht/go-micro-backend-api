package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrSpotNameRequired = errors.New("spot name is required")
	ErrSpotNumberInvalid = errors.New("spot number invalid")
	ErrSpotNameCharMin = errors.New("spot name must have at least 2 characters")
	ErrSportNameFormatInit = errors.New("spot name must start with a capital letter")
	ErrSportNameFormatEnd = errors.New("spot name must end with a number")
	ErrSpotEventIDNotFount = errors.New("spot not found in event")
	ErrorSpotNotFound = errors.New("spot not found")
	ErrorSpotAlreadyReserved = errors.New("spot is already reserved")
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusReserved  SpotStatus = "reserved"
	SpotStatusSold      SpotStatus = "sold"
)

type Spot struct {
	ID         string     `json:"id"`
	EventID    string     `json:"event_id"`
	Name       string     `json:"name"`
	SpotStatus SpotStatus `json:"spot_status"`
	TicketID   string     `json:"ticket_id"`
}


//Validate checks if the spot is valid: Method
func (s Spot) Validate() error {
	if len(s.Name) == 0 {
	 return ErrSpotNameRequired
	}
	if len(s.Name) < 2 {
		return ErrSpotNameCharMin
	}
	//Validade format of name exemple: "A1"
	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSportNameFormatInit
	}
	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSportNameFormatEnd
	}
	return nil
}

//CreatedNewSpot creates a new spot: Function
func CreatedNewSpot (e Event, name string) (*Spot, error) {
	spot := &Spot{
		ID: uuid.New().String(),
		EventID: e.ID,
		Name: name,
		SpotStatus: SpotStatusAvailable,
	}

	if err := spot.Validate(); err != nil {
		return nil, err
	}

	return spot, nil
}


//ReserveSpot reserves a spot: Method
func (s *Spot) ReserveSpot(ticketID string) error {
	if s.SpotStatus == SpotStatusReserved {
		return ErrorSpotAlreadyReserved
	}
	s.SpotStatus = SpotStatusReserved
	s.TicketID = ticketID
	return nil
}
