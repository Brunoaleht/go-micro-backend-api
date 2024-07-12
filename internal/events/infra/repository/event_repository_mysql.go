package repository

import (
	"database/sql"
	"errors"
	"go-backend-api/internal/events/domain"
	"time"
)

type mysqlEventRepository struct {
	db *sql.DB
}

func NewMysqlEventRepository(db *sql.DB) (domain.EventRepository, error) {
	return &mysqlEventRepository{db: db}, nil
}

func (r *mysqlEventRepository) ListEvents() ([]domain.Event, error) {
	query := `SELECT 
	 e.id, e.name, e.location, e.organization,
	 e.rating, e.date, e.image_url, e.capacity, e.price, e.partner_id,
	 s.id, s.event_id, s.name, s.status, s.ticket_id,
	 t.id, t.event_id, t.spot_id, t.ticket_kind, t.price
	 FROM events e
	 LEFT JOIN spots s ON e.id = s.event_id
	 LEFT JOIN tickets t ON s.id = t.spot_id`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	eventMap := make(map[string]*domain.Event)
	spotMap := make(map[string]*domain.Spot)
	for rows.Next() {
		var eventID, eventName, eventLocation, eventOrganization, eventRating, eventImageURL, spotID, spotEventID, spotName, spotStatus, spotTicketID, ticketID, ticketEventID, ticketSpotID, ticketKind sql.NullString
		var eventDate sql.NullString
		var eventCapacity int
		var eventPrice, ticketPrice sql.NullFloat64
		var partnerID sql.NullInt32


		err := rows.Scan(&eventID, &eventName, &eventLocation, 
			&eventOrganization, &eventRating, &eventDate, 
			&eventImageURL, &eventCapacity, &eventPrice, 
			&partnerID, &spotID, &spotEventID, &spotName, 
			&spotStatus, &spotTicketID, &ticketID, &ticketEventID, &ticketSpotID, 
			&ticketKind, &ticketPrice,
		)
		if err != nil {
			return nil, err
		}

		if !eventID.Valid || !eventName.Valid || !eventLocation.Valid || !eventOrganization.Valid || !eventRating.Valid || !eventDate.Valid || !eventImageURL.Valid || !eventPrice.Valid || !partnerID.Valid {
			continue
		}

		event, exists := eventMap[eventID.String]
		if !exists {
			eventDateParsed, err := time.Parse("2006-01-02 15:04:05", eventDate.String)
			if err != nil {
				return nil, err
			}
			event = &domain.Event{
				ID: eventID.String,
				Name: eventName.String,
				Location: eventLocation.String,
				Organization: eventOrganization.String,
				Rating: domain.Rating(eventRating.String),
				Date: eventDateParsed,
				ImageURL: eventImageURL.String,
				Capacity: eventCapacity,
				Price: eventPrice.Float64,
				PartnerID: int(partnerID.Int32),
				Spots: []domain.Spot{},
				Tickets: []domain.Ticket{},
			}
			eventMap[eventID.String] = event
		}

		if spotID.Valid {
			spot, exists := spotMap[spotID.String]
			if !exists {
				spot = &domain.Spot{
					ID: spotID.String,
					EventID: spotEventID.String,
					Name: spotName.String,
					SpotStatus: domain.SpotStatus(spotStatus.String),
					TicketID: spotTicketID.String,
				}
				event.Spots = append(event.Spots, *spot)
				spotMap[spotID.String] = spot
			}

			if ticketID.Valid {
				ticket := domain.Ticket{
					ID: ticketID.String,
					EventID: ticketEventID.String,
					Spot: spot,
					TicketKind: domain.TicketStatus(ticketKind.String),
					Price: ticketPrice.Float64,
				}
				event.Tickets = append(event.Tickets, ticket)
			}
		}
	}

	events := make([]domain.Event, 0, len(eventMap))
	for _, event := range eventMap {
		events = append(events, *event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func (r *mysqlEventRepository) GetEventByID(eventID string) (*domain.Event, error) {
	query := `SELECT 
	 e.id, e.name, e.location, e.organization,
	 e.rating, e.date, e.image_url, e.capacity, e.price, e.partner_id,
	 s.id, s.event_id, s.name, s.status, s.ticket_id,
	 t.id, t.event_id, t.spot_id, t.ticket_kind, t.price
	 FROM events e
	 LEFT JOIN spots s ON e.id = s.event_id
	 LEFT JOIN tickets t ON s.id = t.spot_id
	 WHERE e.id = ?`

	rows, err := r.db.Query(query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var event *domain.Event
	for rows.Next() {
		var eventID, eventName, eventLocation, eventOrganization, eventRating, eventImageURL, spotID, spotEventID, spotName, spotStatus, spotTicketID, ticketID, ticketEventID, ticketSpotID, ticketKind sql.NullString
		var eventDate sql.NullString
		var eventCapacity int
		var eventPrice, ticketPrice sql.NullFloat64
		var partnerID sql.NullInt32

		err := rows.Scan(&eventID, &eventName, &eventLocation, 
			&eventOrganization, &eventRating, &eventDate, 
			&eventImageURL, &eventCapacity, &eventPrice, 
			&partnerID, &spotID, &spotEventID, &spotName, 
			&spotStatus, &spotTicketID, &ticketID, &ticketEventID, &ticketSpotID, 
			&ticketKind, &ticketPrice,
		)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, domain.ErrEventNotFound
			}
			return nil, err
		}

		if !eventID.Valid || !eventName.Valid || !eventLocation.Valid || !eventOrganization.Valid || !eventRating.Valid || !eventDate.Valid || !eventImageURL.Valid || !eventPrice.Valid || !partnerID.Valid {
			continue
		}

		if event == nil {
			eventDateParsed, err := time.Parse("2006-01-02 15:04:05", eventDate.String)
			if err != nil {
				return nil, err
			}
			event = &domain.Event{
				ID: eventID.String,
				Name: eventName.String,
				Location: eventLocation.String,
				Organization: eventOrganization.String,
				Rating: domain.Rating(eventRating.String),
				Date: eventDateParsed,
				ImageURL: eventImageURL.String,
				Capacity: eventCapacity,
				Price: eventPrice.Float64,
				PartnerID: int(partnerID.Int32),
				Spots: []domain.Spot{},
				Tickets: []domain.Ticket{},
			}
		}

		if spotID.Valid {
			spot := domain.Spot{
				ID: spotID.String,
				EventID: spotEventID.String,
				Name: spotName.String,
				SpotStatus: domain.SpotStatus(spotStatus.String),
				TicketID: spotTicketID.String,
			}
			event.Spots = append(event.Spots, spot)

			if ticketID.Valid {
				ticket := domain.Ticket{
					ID: ticketID.String,
					EventID: ticketEventID.String,
					Spot: &spot,
					TicketKind: domain.TicketStatus(ticketKind.String),
					Price: ticketPrice.Float64,
				}
				event.Tickets = append(event.Tickets, ticket)
			}
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if event == nil {
		return nil, domain.ErrEventNotFound
	}

	return event, nil
}

func (r *mysqlEventRepository) CreateEvent(event *domain.Event) error {
	query := `
	INSERT INTO events (id, name, location, organization, rating, date, image_url, capacity, price, partner_id) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query, event.ID, event.Name, event.Location, event.Organization, event.Rating, event.Date.Format("2024-07-15 15:04:05"), event.ImageURL, event.Capacity, event.Price, event.PartnerID)
	if err != nil {
		return err
	}
	return nil
}

func (r *mysqlEventRepository) ReserveSpot(spotID, ticketID string) error {
	query := `UPDATE spots SET spot_status = ?,ticket_id = ? WHERE id = ?`
	_, err := r.db.Exec(query, domain.SpotStatusReserved, ticketID, spotID)
	if err != nil {
		return err
	}
	return nil
}

func (r *mysqlEventRepository) CreateSpot(spot *domain.Spot) error {
	query := `INSERT INTO spots (id, event_id, name, spot_status, ticket_id) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, spot.ID, spot.EventID, spot.Name, spot.SpotStatus, spot.TicketID)
	if err != nil {
		return err
	}
	return nil
}

func (r *mysqlEventRepository) CreateTicket(ticket *domain.Ticket) error {
	query := `INSERT INTO tickets (id, event_id, spot_id, ticket_kind, price) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, ticket.ID, ticket.EventID, ticket.Spot.ID, ticket.TicketKind, ticket.Price)
	if err != nil {
		return err
	}
	return nil
}

func (r *mysqlEventRepository) FindSpotsEventID(eventID string) ([]domain.Spot, error) {
	query := `
		SELECT id, event_id, name, spot_status, ticket_id
		FROM spots
		WHERE event_id = ?
	`
	rows, err := r.db.Query(query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spots []domain.Spot
	for rows.Next() {
		var spot domain.Spot
		err := rows.Scan(&spot.ID, &spot.EventID, &spot.Name, &spot.SpotStatus, &spot.TicketID)
		if err != nil {
			return nil, err
		}
		spots = append(spots, spot)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return spots, nil
}

func (r *mysqlEventRepository) FindSpotById(spotID string) (*domain.Spot, error) {
	query := `
		SELECT id, event_id, name, spot_status, ticket_id
		FROM spots
		WHERE id = ?
	`
	row := r.db.QueryRow(query, spotID)
	var spot domain.Spot
	err := row.Scan(&spot.ID, &spot.EventID, &spot.Name, &spot.SpotStatus, &spot.TicketID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrorSpotNotFound
		}
		return nil, err
	}

	return &spot, nil
}

func (r *mysqlEventRepository) FindSpotByName(eventID, spotName string) (*domain.Spot, error) {
	query := `
		SELECT 
			s.id, s.event_id, s.name, s.status, s.ticket_id,
			t.id, t.event_id, t.spot_id, t.ticket_kind, t.price
		FROM spots s
		LEFT JOIN tickets t ON s.id = t.spot_id
		WHERE s.event_id = ? AND s.name = ?
	`
	rows, err := r.db.Query(query, eventID, spotName)
	if err != nil {
		return nil, err
	}

	var spot domain.Spot
	var ticket domain.Ticket
	var ticketID, ticketEventID, ticketSpotID, ticketKind sql.NullString
	var ticketPrice sql.NullFloat64

	for rows.Next() {
		err := rows.Scan(&spot.ID, &spot.EventID, &spot.Name, &spot.SpotStatus, &spot.TicketID, &ticketID, &ticketEventID, &ticketSpotID, &ticketKind, &ticketPrice)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, domain.ErrorSpotNotFound
			}
			return nil, err
		}

		if ticketID.Valid {
			ticket = domain.Ticket{
				ID: ticketID.String,
				EventID: ticketEventID.String,
				Spot: &spot,
				TicketKind: domain.TicketStatus(ticketKind.String),
				Price: ticketPrice.Float64,
			}
			spot.TicketID = ticket.ID
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &spot, nil
}