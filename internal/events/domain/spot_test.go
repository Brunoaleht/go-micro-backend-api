package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreatedNewSpot(t *testing.T) {
	event, err := CreatedNewEvent("Event Test", "Location Test", "Organization Test", RatingFree, time.Now().Add(24*time.Hour), "image_url", 100, 50.00, 1)
	assert.Nil(t, err)
	assert.NotNil(t, event)
	
	spot, err := CreatedNewSpot(*event, "A1")
	assert.Nil(t, err)
	assert.NotNil(t, spot)
	assert.Equal(t, "A1", spot.Name)
	assert.Equal(t, event.ID, spot.EventID)
	assert.Equal(t, SpotStatusAvailable, spot.SpotStatus)
	assert.Empty(t, spot.TicketID)
	assert.NotEmpty(t, spot.ID)
}

func TestSpot_Validate(t *testing.T) {
	spot := Spot{
		Name: "",
	}
	err := spot.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrSpotNameRequired, err)

	spot.Name = "A"
	err = spot.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrSpotNameCharMin, err)

	spot.Name = "1"
	err = spot.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrSpotNameCharMin, err)

	spot.Name = "a1"
	err = spot.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrSportNameFormatInit, err)

	spot.Name = "Aa"
	err = spot.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrSportNameFormatEnd, err)

	spot.Name = "A1"
	err = spot.Validate()
	assert.Nil(t, err)
}

func TestSpot_Reserve(t *testing.T) {
	event, _ := CreatedNewEvent("Event Test", "Location Test", "Organization Test", RatingFree, time.Now().Add(24*time.Hour), "image_url", 100, 50.00, 1)
	assert.NotNil(t, event)

	spot, _ := CreatedNewSpot(*event, "A1")
	assert.NotNil(t, spot)

	 err := spot.ReserveSpot("Ticket123")
	 assert.Nil(t, err)
	 assert.Equal(t, SpotStatusReserved, spot.SpotStatus)
	 assert.Equal(t, "Ticket123", spot.TicketID)
}