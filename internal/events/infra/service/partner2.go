package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Partner2 struct {
	BaseURL string
}

type Partner2ReservationRequest struct {
	Spots			[]string `json:"spots"`
	TicketKind		string `json:"ticket_kind"`
	Email			string `json:"email"`
}

type Partner2ReservationResponse struct {
	ID string `json:"id"`
	Email string `json:"email"`
	Spot string `json:"spot"`
	TicketKind string `json:"ticket_kind"`
	Status string `json:"status"`
	EventID string `json:"event_id"`
}

func (p *Partner2) MakeReservation(req *ReservationRequest) ([]ReservationResponse, error) {
	TicketKind := req.TicketKind
	if TicketKind == "full" {
		TicketKind = "full"
	} else {
		TicketKind = "half"
	}

	partnerReq := Partner2ReservationRequest{
		Spots: req.Spots,
		TicketKind: TicketKind,
		Email: req.Email,
	}

	body, err := json.Marshal(partnerReq)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/matters/%s/reserve", p.BaseURL, req.EventID)
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("reservation failed with status code: %d", httpResp.StatusCode)
	}

	var partnerResp []Partner2ReservationResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&partnerResp); err != nil {
		return nil, err
	}

	responses := make([]ReservationResponse, len(partnerResp))
	for i, resp := range partnerResp {
		responses[i] = ReservationResponse{
			ID: resp.ID,
			Spot: resp.Spot,
			Status: resp.Status,
		
		}
	}

	return responses, nil
}