package usecase

import (
	"go-backend-api/internal/events/domain"
)

type GetEventInputDto struct {
	ID string `json:"id"`
}

type GetEventOutputDto struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Location     string  `json:"location"`
	Organization string  `json:"organization"`
	Rating       string  `json:"rating"`
	Date         string  `json:"date"`
	Capacity     int     `json:"capacity"`
	Price        float64 `json:"price"`
	PartnerID    int     `json:"partner_id"`
}

type GetEventUseCase struct {
	repo domain.EventRepository
}

func NewGetEventUseCase(repo domain.EventRepository) *GetEventUseCase {
	return &GetEventUseCase{repo: repo}
}

func (uc *GetEventUseCase) Execute(input GetEventInputDto) (*GetEventOutputDto, error) {
	event, err := uc.repo.GetEventByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetEventOutputDto{
		ID:           event.ID,
		Name:         event.Name,
		Location:     event.Location,
		Organization: event.Organization,
		Rating:       string(event.Rating),
		Date:         event.Date.Format("2006-01-02 15:04:05"),
		Capacity:     event.Capacity,
		Price:        event.Price,
		PartnerID:    event.PartnerID,
	}, nil
}