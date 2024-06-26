package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewTicket(t *testing.T) {
	event, err := CreatedNewEvent("Event Test", "Location Test", "Organization Test", RatingFree, time.Now().Add(24*time.Hour), "image_url", 100, 50.00, 1)
	assert.Nil(t, err)
	assert.NotNil(t, event)

	spot, err := CreatedNewSpot(*event, "A1")
	assert.Nil(t, err)
	assert.NotNil(t, spot)

	ticket, err := CreatedNewTicket(event, spot,  TicketStatusFull)
	assert.Nil(t, err)
	assert.NotNil(t, ticket)
	assert.Equal(t, event.ID, ticket.EventID)
	assert.Equal(t, spot.ID, ticket.Spot.ID)
	assert.Equal(t, TicketStatusFull, ticket.TicketKind)
	assert.Equal(t, 50.00, ticket.Price)
	assert.NotEmpty(t, ticket.ID)
}

func TestTicket_CalculatePrice(t *testing.T) {
	event, err := CreatedNewEvent("Event Test", "Location Test", "Organization Test", RatingFree, time.Now().Add(24*time.Hour), "image_url", 100, 50.00, 1)
	assert.Nil(t, err)
	assert.NotNil(t, event)

	spot, err := CreatedNewSpot(*event, "A1")
	assert.Nil(t, err)
	assert.NotNil(t, spot)

	ticket, err := CreatedNewTicket(event, spot,  TicketStatusHalf)
	assert.Nil(t, err)
	assert.NotNil(t, ticket)
	assert.Equal(t, 25.00, ticket.Price)
}

func TestTicket_Validate(t *testing.T) {
	event, err := CreatedNewEvent("Event Test", "Location Test", "Organization Test", RatingFree, time.Now().Add(24*time.Hour), "image_url", 100, 50.00, 1)
	assert.Nil(t, err)
	assert.NotNil(t, event)

	spot, err := CreatedNewSpot(*event, "A1")
	assert.Nil(t, err)
	assert.NotNil(t, spot)

	ticket := Ticket{
		Spot: nil,
		Price: 50.00,
	}
	err = ticket.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrTicketSpotRequired, err)

	ticket.Spot = spot
	ticket.Price = 0
	err = ticket.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrTicketPriceInvalid, err)
}