package http

import (
	"encoding/json"
	"go-backend-api/internal/events/usecase"
	"net/http"
)

// EventsHandler handles HTTP the events requests
type EventsHandler struct {
	listEventsUseCase *usecase.ListEventsUseCase
	getEventUseCase  *usecase.GetEventUseCase
	createEventUseCase *usecase.CreateEventUseCase
	buyTicketsUseCase	*usecase.BuyTicketsUseCase
	createSpotsUseCase	*usecase.CreateSpotsUseCase
	listSpotsUseCase	*usecase.ListSpotsUseCase
}

// NewEventsHandler creates a new EventsHandler
func NewEventsHandler(
	listEventsUseCase *usecase.ListEventsUseCase,
	getEventUseCase *usecase.GetEventUseCase,
	createEventUseCase *usecase.CreateEventUseCase,
	buyTicketsUseCase *usecase.BuyTicketsUseCase,
	createSpotsUseCase *usecase.CreateSpotsUseCase,
	listSpotsUseCase *usecase.ListSpotsUseCase,
) *EventsHandler {
	return &EventsHandler{
		listEventsUseCase: listEventsUseCase,
		getEventUseCase: getEventUseCase,
		createEventUseCase: createEventUseCase,
		buyTicketsUseCase: buyTicketsUseCase,
		createSpotsUseCase: createSpotsUseCase,
		listSpotsUseCase: listSpotsUseCase,
	}
}

// ListEvents handles the request to list all events.
// @Summary List all events
// @Description Get all events with their details
// @Tags Events
// @Accept json
// @Produce json
// @Success 200 {object} usecase.ListEventsOutputDTO
// @Failure 500 {object} string
// @Router /events [get]
func (h *EventsHandler) ListEvents(w http.ResponseWriter, r *http.Request) {
	output, err := h.listEventsUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

// GetEvent handles the request to get an event by its ID.
// @Summary Get an event
// @Description Get an event by its ID
// @Tags Events
// @Accept json
// @Produce json
// @Param eventId path string true "Event ID"
// @Success 200 {object} usecase.GetEventOutputDTO
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /events/{eventId} [get]
func (h *EventsHandler) GetEvent(w http.ResponseWriter, r *http.Request){
	eventID := r.PathValue("eventId")
	input := usecase.GetEventInputDto{ID: eventID}

	output, err := h.getEventUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

// CreateEvent handles the request to create a new event.
// @Summary Create an event
// @Description Create a new event
// @Tags Events
// @Accept json
// @Produce json
// @Param event body usecase.CreateEventInputDto true "Event data"
// @Success 201 {object} usecase.CreateEventOutputDto
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *EventsHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateEventInputDto
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.createEventUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

// BuyTickets handles the request to buy tickets for an event.
// @Summary Buy tickets
// @Description Buy tickets for an event
// @Tags Events
// @Accept json
// @Produce json
// @Param body usecase.BuyTicketsInputDto true "Tickets data"
// @Success 201 {object} usecase.BuyTicketsOutputDto
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /events/buy-tickets [post]
func (h *EventsHandler) BuyTickets(w http.ResponseWriter, r *http.Request) {
	var input usecase.BuyTicketsInputDto
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.buyTicketsUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

// CreateSpots handles the request to create spots for an event.
// @Summary Create spots
// @Description Create spots for an event
// @Tags Events
// @Accept json
// @Produce json
// @Param eventId path string true "Event ID"
// @Param body usecase.CreateSpotsInputDto true "Spots data"
// @Success 201 {object} usecase.CreateSpotsOutputDto
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /events/{eventId}/spots [post]
func (h *EventsHandler) CreateSpots(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventId")
	var input usecase.CreateSpotsInputDto
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	input.EventID = eventID

	output, err := h.createSpotsUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

//WriteErrorResponse writes an error response json format
func (h *EventsHandler) WriteErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Message string `json:"message"`
}

//CreateSpotsRequest represents the request to create spots
type CreateSpotsRequest struct {
	NumberOfSpots int `json:"number_of_spots"`
}

//ListSpots handles the request to list all spots for an event.
func(h *EventsHandler) ListSpots(w http.ResponseWriter, r *http.Request){
	eventID := r.PathValue("eventId")
	input := usecase.ListSpotsInputDto{EventID: eventID}

	output, err := h.listSpotsUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}