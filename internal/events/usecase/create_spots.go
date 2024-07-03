package usecase

import (
	"go-backend-api/internal/events/domain"

	"fmt"
)

type CreateSpotsInputDto struct {
	EventID string `json:"event_id"`
	NumberOfSpots int `json:"number_of_spots"`
}

type CreateSpotsOutputDto struct {
	Spots []SpotDto `json:"spots"`
}

type CreateSpotsUseCase struct {
	repo domain.EventRepository
}

func NewCreateSpotsUseCase(repo domain.EventRepository) *CreateSpotsUseCase {
	return &CreateSpotsUseCase{repo: repo}
}

func (uc *CreateSpotsUseCase) Execute(input CreateSpotsInputDto) (*CreateSpotsOutputDto, error) {
	event, err := uc.repo.GetEventByID(input.EventID)
	if err != nil {
		return nil, err
	}

	spots := make([]domain.Spot, input.NumberOfSpots)
	for i := 0; i < input.NumberOfSpots; i++ {
		spotName := generateSpotName(i)
		spot, err := domain.CreatedNewSpot(*event, spotName)
		if err != nil {
			return nil, err
		}
		if err := uc.repo.CreateSpot(spot); err != nil {
			return nil, err
		}
		spots[i] = *spot
	}

	spotDto := make([]SpotDto, len(spots))
	for i, spot := range spots {
		spotDto[i] = SpotDto{
			ID:         spot.ID,
			EventID:    spot.EventID,
			Name:       spot.Name,
			SpotStatus: string(spot.SpotStatus),
			TicketID:   spot.TicketID,
		}
	}

	return &CreateSpotsOutputDto{Spots: spotDto}, nil
}

func generateSpotName(index int) string {
	letter := 'A' + rune(index/10)
	number := index%10 + 1
	return fmt.Sprintf("%c-%d", letter, number)
}
