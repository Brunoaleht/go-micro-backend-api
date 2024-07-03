package usecase

import (
	"time"

	"go-backend-api/internal/events/domain"
)

type CreateEventInputDto struct {
	Name         string    `json:"name"`
	Location     string    `json:"location"`
	Organization string    `json:"organization"`
	Rating       string    `json:"rating"`
	Date         time.Time `json:"date"`
	Capacity     int       `json:"capacity"`
	ImageURL     string  `json:"image_url"`
	Price        float64   `json:"price"`
	PartnerID    int       `json:"partner_id"`
}

type CreateEventOutputDto struct {
	ID string `json:"id"`
	Name         string    `json:"name"`
	Location     string    `json:"location"`
	Organization string    `json:"organization"`
	Rating       string    `json:"rating"`
	Date         time.Time `json:"date"`
	Capacity     int       `json:"capacity"`
	ImageURL     string  `json:"image_url"`
	Price        float64   `json:"price"`
	PartnerID    int       `json:"partner_id"`
}

type CreateEventUseCase struct {
	repo domain.EventRepository
}

func NewCreateEventUseCase(repo domain.EventRepository) *CreateEventUseCase {
	return &CreateEventUseCase{repo: repo}
}

func (uc *CreateEventUseCase) Execute(input CreateEventInputDto) (*CreateEventOutputDto, error) {
	event, err := domain.CreatedNewEvent(
		input.Name, 
		input.Location, 
		input.Organization, 
		domain.Rating(input.Rating), 
		input.Date, 
		input.ImageURL, 
		input.Capacity, 
		input.Price, 
		input.PartnerID,
	)
	if err != nil {
		return &CreateEventOutputDto{}, err
	}

	err = uc.repo.CreateEvent(event)
	if err != nil {
		return &CreateEventOutputDto{}, err
	}

	return &CreateEventOutputDto{
		ID:           event.ID,
		Name:         event.Name,
		Location:     event.Location,
		Organization: event.Organization,
		Rating:       string(event.Rating),
		Date:         event.Date,
		Capacity:     event.Capacity,
		ImageURL:     event.ImageURL,
		Price:        event.Price,
		PartnerID:    event.PartnerID,
	}, nil
}