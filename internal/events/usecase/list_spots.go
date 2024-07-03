package usecase

import (
	"go-backend-api/internal/events/domain"
)

type SpotDto struct {
	ID         string     `json:"id"`
	EventID    string     `json:"event_id"`
	Name       string     `json:"name"`
	SpotStatus string `json:"spot_status"`
	TicketID   string     `json:"ticket_id"`
	Reserved bool   `json:"reserved"`
}

type ListSpotsInputDto struct {
	EventID string `json:"event_id"`
}

type ListSpotsOutputDto struct {
	Spots []SpotDto `json:"spots"`
	Event EventDto `json:"event"`
}

type ListSpotsUseCase struct {
	repo domain.EventRepository
}

func NewListSpotsUseCase(repo domain.EventRepository) *ListSpotsUseCase {
	return &ListSpotsUseCase{repo: repo}
}

func (uc *ListSpotsUseCase) Execute(input ListSpotsInputDto) (*ListSpotsOutputDto, error) {
	event, err := uc.repo.GetEventByID(input.EventID)
	if err != nil {
		return nil, err
	}

	spots, err := uc.repo.FindSpotsEventID(input.EventID)
	if err != nil {
		return nil, err
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

	eventDto := EventDto{
		ID:           event.ID,
		Name:         event.Name,
		Location:     event.Location,
		Organization: event.Organization,
		Rating:       string(event.Rating),
		Date:         event.Date.Format("2006-01-02 15:04:05"),
		ImageURL:     event.ImageURL,
		Capacity:     event.Capacity,
		Price:        event.Price,
		PartnerID:    event.PartnerID,
	}

	return &ListSpotsOutputDto{Spots: spotDto, Event: eventDto}, nil

}