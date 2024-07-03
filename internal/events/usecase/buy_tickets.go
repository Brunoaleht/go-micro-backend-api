package usecase

import (
	"go-backend-api/internal/events/domain"
)

type BuyTicketsInputDto struct {
	EventID string `json:"event_id"`
	Spots []string `json:"spots"`
	TicketKind string `json:"ticket_status"`
	CardHash string `json:"card_hash"`
	Email string `json:"email"`
}

type BuyTicketsOutputDto struct {
	Tickets []TicketDto `json:"tickets"`
}

type TicketDto struct {
	ID string `json:"id"`
	SpotID string `json:"spot_id"`
	EventID string `json:"event_id"`
	TicketKind string `json:"ticket_status"`
	Price float64 `json:"price"`
}

type BuyTicketsUseCase struct {
	repo domain.EventRepository
	partnerFactory service.PartnerFactory
}

func NewBuyTicketsUseCase(repo domain.EventRepository, partnerFactory service.PartnerFactory) *BuyTicketsUseCase {
	return &BuyTicketsUseCase{
		repo: repo, 
		partnerFactory: partnerFactory,
	}
}

func (uc *BuyTicketsUseCase) Execute(input BuyTicketsInputDto) (*BuyTicketsOutputDto, error) {
	event, err := uc.repo.GetEventByID(input.EventID)
	if err != nil {
		return nil, err
	}

	//criar a solicitação de reserva
	reserver := &service.ReservationRequest{
		EventID: input.EventID,
		Spots: input.Spots,
		TicketKind: input.TicketKind,
		CardHash: input.CardHash,
		Email: input.Email,
		
	}

	//Obtendo o parceiro
	partnerService, err := uc.partnerFactory.GetPartner(event.PartnerID)
	if err != nil {
		return nil, err
	}

	//Reservar os tickets usando o serviço do parceiro
	reservationResponse, err := partnerService.MakeReservation(reserver)
	if err != nil {
		return nil, err
	}

	//salvando os tickets no banco de dados
	tickets := make([]domain.Ticket, len(reservationResponse))
	for i, reservation := range reservationResponse {
		spot, err := uc.repo.FindSpotByName(event.ID, reservation.Spot)
		if err != nil {
			return nil, err
		}

		ticket, err := domain.CreatedNewTicket(event, spot, domain.TicketStatus(input.TicketKind))
		if err != nil {
			return nil, err
		}

		err = uc.repo.CreateTicket(ticket)
		if err != nil {
			return nil, err
		}

		spot.ReserveSpot(ticket.ID)
		err = uc.repo.ReserveSpot(spot.ID, ticket.ID)
		if err != nil {
			return nil, err
		}

		tickets[i] = *ticket
	}

	ticketDto := make([]TicketDto, len(tickets))
	for i, ticket := range tickets {
		ticketDto[i] = TicketDto{
			ID: ticket.ID,
			SpotID: ticket.Spot.ID,
			TicketKind: string(ticket.TicketKind),
			Price: ticket.Price,
		}
	}
	
	return &BuyTicketsOutputDto{Tickets: ticketDto}, nil
}