package domain

import (
	"errors"
	"fmt"
)

type spotService struct{}

var (
	ErrInvalidNumberSpot = errors.New("number of spot must be greater than zero")
)

// NewSpotService creates a new spot service
func NewSpotService() *spotService {
	return &spotService{}
}

func (s *spotService) GenerateSpots(event *Event, numberSpot int) error {
	if numberSpot <= 0 {
		return ErrInvalidNumberSpot
	}

	for i := 0; i < numberSpot; i++ {
		spotName := fmt.Sprintf("%c%d", 'A'+i, i%10+1)
		spot, err := CreatedNewSpot(*event, spotName)
		if err != nil {
			return err
		}
		event.Spots = append(event.Spots, *spot)
	}
	return nil

}