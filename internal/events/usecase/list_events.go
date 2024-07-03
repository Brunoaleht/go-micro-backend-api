package usecase

import "go-backend-api/internal/events/domain"

type ListEventsOutputDto struct {
	Events []EventDto `json:"events"`
}

type EventDto struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Location     string  `json:"location"`
	Organization string  `json:"organization"`
	Rating       string  `json:"rating"`
	Date         string  `json:"date"`
	ImageURL     string  `json:"image_url"`
	Capacity     int     `json:"capacity"`
	Price        float64 `json:"price"`
	PartnerID    int     `json:"partner_id"`
}

type ListEventsUseCase struct {
	repo domain.EventRepository
}

func NewListEventsUseCase(repo domain.EventRepository) *ListEventsUseCase {
	return &ListEventsUseCase{repo: repo}
}

func (u *ListEventsUseCase) Execute() (*ListEventsOutputDto, error) {
	events, err := u.repo.ListEvents()
	if err != nil {
		return nil, err
	}

	eventDto := make([]EventDto, len(events))
	for i, event := range events {
		eventDto[i] = EventDto{
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
	}

	return &ListEventsOutputDto{Events: eventDto}, nil
}