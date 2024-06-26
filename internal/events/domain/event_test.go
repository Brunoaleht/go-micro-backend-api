package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


func TestCreatedNewEvent(t *testing.T) {
	event, err := CreatedNewEvent("Event Test", "Location Test", "Organization Test", RatingFree, time.Now().Add(24*time.Hour), "image_url", 100, 50.00, 1)
	assert.Nil(t, err)
	assert.NotNil(t, event)
	assert.Equal(t, "Event Test", event.Name)
	assert.Equal(t, "Location Test", event.Location)
	assert.Equal(t, "Organization Test", event.Organization)
	assert.Equal(t, RatingFree, event.Rating)
	assert.Equal(t, 100, event.Capacity)
	assert.Equal(t, 50.00, event.Price)
	assert.Equal(t, 1, event.PartnerID)
	assert.NotEmpty(t, event.ID)
	assert.Empty(t, event.Spots)
	assert.Empty(t, event.Tickets)
}

func TestEvent_Validate(t *testing.T) {
	event := Event{
		Name: "",
		Date: time.Now().Add(24 * time.Hour),
		Capacity: 100,
		Price: 50.00,
	}

	err := event.Validade()
	assert.NotNil(t, err)
	assert.Equal(t, ErrEventNameRequired, err)

	event.Name = "Event Test"
	event.Date = time.Now().Add(-24 * time.Hour)
	err = event.Validade()
	assert.NotNil(t, err)
	assert.Equal(t, ErrEventDateInFuture, err)

	event.Date = time.Now().Add(24 * time.Hour)
	event.Capacity = 0
	err = event.Validade()
	assert.NotNil(t, err)
	assert.Equal(t, ErrEventCapacityInvalid, err)

	event.Capacity = 100
	event.Price = 0
	err = event.Validade()
	assert.NotNil(t, err)
	assert.Equal(t, ErrEventPriceInvalid, err)
}

func TestEvent_AddSpot(t *testing.T) {
	event, err := CreatedNewEvent("Event Test", "Location Test", "Organization Test", RatingFree, time.Now().Add(24*time.Hour), "image_url", 100, 50.00, 1)
	assert.Nil(t, err)
	assert.NotNil(t, event)

	spot, err := event.AddSpot("A1")
	assert.Nil(t, err)
	assert.NotNil(t, spot)
	assert.Equal(t, "A1", spot.Name)
	assert.Equal(t, event.ID, spot.EventID)
	assert.Equal(t, SpotStatusAvailable, spot.SpotStatus)
	assert.Equal(t, 1, len(event.Spots))
}