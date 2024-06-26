package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEventNameRequired = errors.New("event name is required")
	ErrEventDateInFuture = errors.New("event date must be in the future")
	ErrEventCapacityInvalid = errors.New("event capacity must be greater than zero")
	ErrEventPriceInvalid = errors.New("event price must be greater than zero")
)


type Rating string

const (
	RatingFree 	Rating = "L"
	Rating10 	Rating = "L10"
	Rating12 	Rating = "L12"
	Rating14 	Rating = "L14"
	Rating16 	Rating = "L16"
	Rating18 	Rating = "L18"
)

type Event struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	Organization string `json:"organization"`
	Rating       Rating    `json:"rating"`
	Date         time.Time `json:"date"`
	ImageURL     string `json:"image_url"`
	Capacity     int    `json:"capacity"`
	Price        float64    `json:"price"`
	PartnerID    int `json:"partner_id"`
	Spots        []Spot `json:"spots"`
	Tickets			 []Ticket `json:"tickets"`
}

func CreatedNewEvent(name, location, organization string, rating Rating, date time.Time, imageURL string, capacity int, price float64, partnerID int) (*Event, error) {
	event := &Event{
		ID:           uuid.New().String(),
		Name:         name,
		Location:     location,
		Organization: organization,
		Rating:       rating,
		Date:         date,
		ImageURL:     imageURL,
		Capacity:     capacity,
		Price:        price,
		PartnerID:    partnerID,
		Spots: 			make([]Spot, 0),
	}
	if err := event.Validade(); err != nil {
		return nil, err
	}
	return event, nil
}

func (e Event) Validade() error {
	if e.Name == "" {
		return ErrEventNameRequired
	}
	if e.Date.Before(time.Now()) {
		return ErrEventDateInFuture
	}
	if e.Capacity <= 0 {
		return ErrEventCapacityInvalid
	}
	if e.Price <=0 {
		return ErrEventPriceInvalid
	}
	return nil
} 

func (e *Event) AddSpot(name string) (*Spot, error) {
		spot, err := CreatedNewSpot(*e, name)
		if err != nil {
			return nil, err
		}

		e.Spots = append(e.Spots, *spot)
		return spot, nil
}