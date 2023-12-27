package eventserver

import (
	"encoding/json"
	"io"
	"time"

	"wb_l2/internal/domain"
)

type createEventRequest struct {
	Title string    `json:"title,omitempty"`
	Date  time.Time `json:"date,omitempty"`
}

// GetEvent generates a new Event structure from the requests
func (r *createEventRequest) GetEvent() (output *domain.Event) {
	return &domain.Event{
		Title: r.Title,
		Date:  r.Date,
	}
}

// decodePOSTRequest универальная функция для всех POST запросов
func decodePOSTRequest(r io.Reader, target any) (err error) {
	err = json.NewDecoder(r).Decode(target)
	return
}

type deleteEventRequest struct {
	ID string
}

type updateEventRequest struct {
	ID    string    `json:"id,omitempty"`
	Title string    `json:"title,omitempty"`
	Date  time.Time `json:"date,omitempty"`
}

// GetEvent generates a new Event structure from the request
func (r *updateEventRequest) GetEvent() (output *domain.Event) {
	return &domain.Event{
		ID:    r.ID,
		Title: r.Title,
		Date:  r.Date,
	}
}
